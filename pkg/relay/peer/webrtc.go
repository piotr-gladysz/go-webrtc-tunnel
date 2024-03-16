package peer

import (
	"github.com/pion/webrtc/v4"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/relay/proxymsg"
	"log/slog"
	"slices"
)

func (p *Peer) AddICECandidate(candidateStr string) error {
	candidate := webrtc.ICECandidateInit{Candidate: candidateStr}
	return p.peerConnection.AddICECandidate(candidate)
}

func (p *Peer) SetRemoteDescription(descrStr string) error {
	descr := webrtc.SessionDescription{
		Type: webrtc.SDPTypeAnswer,
		SDP:  descrStr,
	}

	if err := p.peerConnection.SetRemoteDescription(descr); err != nil {
		return err
	}
	if len(p.iceCandidates) > 0 && p.webrtcListener != nil {
		p.candidateMux.Lock()
		defer p.candidateMux.Unlock()

		for _, candidateStr := range p.iceCandidates {
			candidate := webrtc.ICECandidateInit{Candidate: candidateStr}
			if err := p.peerConnection.AddICECandidate(candidate); err != nil {
				slog.Warn("Failed to add ICE candidate", "err", err)
			}
		}
		p.iceCandidates = make([]string, 0)
	}
	return nil
}

func (p *Peer) RecvMessage(msg webrtc.DataChannelMessage) {
	p.proxyRecvMux.Lock()
	defer p.proxyRecvMux.Unlock()

	decoded, err := p.proxyDecoder.Decode(msg.Data)
	if err != nil {
		slog.Warn("Failed to decode message", "err", err)

		// TODO: callback message
		return
	}

	if !slices.Contains(p.AllowedLocalPorts, int(decoded.Port)) {
		slog.Warn("Port not allowed", "port", decoded.Port)

		// TODO: callback message
		return
	}

	if err = p.proxyMessageReceiver.RecvMessage(decoded); err != nil {
		slog.Warn("Failed to read message", "err", err)
	}
}

func (p *Peer) SendMessage(msg proxymsg.ProxyMessage) error {
	p.proxySendMux.Lock()
	defer p.proxySendMux.Unlock()

	buf := make([]byte, msg.GetTotalLen())

	err := p.proxyEncoder.Encode(&msg, buf)
	if err != nil {
		slog.Warn("Failed to encode message", "err", err)
		return err
	}

	return p.controlChannel.Send(buf)
}

func (p *Peer) onICEConnectionStateChange(state webrtc.ICEConnectionState) {
	switch state {
	case webrtc.ICEConnectionStateConnected:
		slog.Debug("ICE state change", "state", "connected")
	case webrtc.ICEConnectionStateDisconnected:
		slog.Debug("ICE state change", "state", "disconnected")
	case webrtc.ICEConnectionStateFailed:
		slog.Debug("ICE state change", "state", "failed")
	case webrtc.ICEConnectionStateClosed:
		slog.Debug("ICE state change", "state", "closed")
	case webrtc.ICEConnectionStateChecking:
		slog.Debug("ICE state change", "state", "checking")
	case webrtc.ICEConnectionStateNew:
		slog.Debug("ICE state change", "state", "new")
	case webrtc.ICEConnectionStateCompleted:
		slog.Debug("ICE state change", "state", "completed")
	default:
		slog.Error("Unknown ICE state", "state", state)
	}

	if p.stateListener != nil {
		p.stateListener.OnICEConnectionStateChange(p, state)
	}
}

func (p *Peer) onDataChannel(channel *webrtc.DataChannel) {
	slog.Debug("On data channel", "peerId", p.Id, "label", channel.Label())
	if p.stateListener != nil {
		p.stateListener.OnDataChannel(p, channel)
	}

	channel.OnMessage(p.RecvMessage)
}

func (p *Peer) onICECandidate(candidate *webrtc.ICECandidate) {
	if candidate == nil {
		slog.Warn("nil ICE candidate")
		return
	}

	candidateStr := candidate.ToJSON().Candidate
	slog.Debug("On ICE candidate", "peerId", p.Id, "ice", candidateStr)

	p.candidateMux.Lock()
	defer p.candidateMux.Unlock()

	if p.peerConnection.RemoteDescription() == nil || p.webrtcListener == nil {
		p.iceCandidates = append(p.iceCandidates, candidateStr)
	} else {
		p.webrtcListener.OnICECandidate(p, candidate)
	}
}
