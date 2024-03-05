package tunnel

import (
	"github.com/pion/webrtc/v4"
	"log/slog"
)

func (t *Tunnel) AddICECandidate(candidateStr string) error {
	candidate := webrtc.ICECandidateInit{Candidate: candidateStr}
	return t.peerConnection.AddICECandidate(candidate)
}

func (t *Tunnel) SetRemoteDescription(descrStr string) error {
	descr := webrtc.SessionDescription{
		Type: webrtc.SDPTypeAnswer,
		SDP:  descrStr,
	}

	if err := t.peerConnection.SetRemoteDescription(descr); err != nil {
		return err
	}
	if len(t.iceCandidates) > 0 && t.listener != nil {
		t.candidateMux.Lock()
		defer t.candidateMux.Unlock()

		for _, candidateStr := range t.iceCandidates {
			candidate := webrtc.ICECandidateInit{Candidate: candidateStr}
			if err := t.peerConnection.AddICECandidate(candidate); err != nil {
				slog.Warn("Failed to add ICE candidate", "err", err)
			}
		}
		t.iceCandidates = make([]string, 0)
	}
	return nil
}

func (t *Tunnel) onICEConnectionStateChange(state webrtc.ICEConnectionState) {
	switch state {
	case webrtc.ICEConnectionStateConnected:
	case webrtc.ICEConnectionStateDisconnected:
	case webrtc.ICEConnectionStateFailed:
	case webrtc.ICEConnectionStateClosed:

	}
}

func (t *Tunnel) onICECandidate(candidate *webrtc.ICECandidate) {
	if candidate == nil {
		slog.Warn("nil ICE candidate")
		return
	}

	candidateStr := candidate.ToJSON().Candidate
	slog.Debug("On ICE candidate", "ice", candidateStr)

	t.candidateMux.Lock()
	defer t.candidateMux.Unlock()

	if t.peerConnection.RemoteDescription() == nil || t.listener == nil {
		t.iceCandidates = append(t.iceCandidates, candidateStr)
	} else {
		t.listener.OnICECandidate(t, candidate)
	}
}
