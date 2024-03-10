package peer

import (
	"github.com/pion/webrtc/v4"
	"log/slog"
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
	if len(p.iceCandidates) > 0 && p.webrtclistener != nil {
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

	if p.stateListener != nil {
		p.stateListener.OnDataChannel(p, channel)
	}
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

	if p.peerConnection.RemoteDescription() == nil || p.webrtclistener == nil {
		p.iceCandidates = append(p.iceCandidates, candidateStr)
	} else {
		p.webrtclistener.OnICECandidate(p, candidate)
	}
}
