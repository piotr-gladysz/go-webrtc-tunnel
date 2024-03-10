package peer

import (
	"context"
	"github.com/pion/webrtc/v4"
	"log/slog"
	"os"
	"testing"
	"time"
)

func init() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	handler := slog.NewTextHandler(os.Stdout, opts)
	slog.SetDefault(slog.New(handler))
}

type WebrtcListenerMock struct {
	onIceCandidate     func(p *Peer, candidate *webrtc.ICECandidate)
	onLocalDescription func(p *Peer, sdp webrtc.SessionDescription)
	onConnected        func(p *Peer)
}

type StateListenerMock struct {
	onICEConnectionStateChange func(p *Peer, state webrtc.ICEConnectionState)
	onDataChannel              func(p *Peer, channel *webrtc.DataChannel)
}

func (p *WebrtcListenerMock) OnICECandidate(t *Peer, candidate *webrtc.ICECandidate) {
	p.onIceCandidate(t, candidate)
}

func (p *WebrtcListenerMock) OnLocalDescription(t *Peer, sdp webrtc.SessionDescription) {
	p.onLocalDescription(t, sdp)
}

func (p *WebrtcListenerMock) OnConnected(t *Peer) {
	p.onConnected(t)
}

func (s *StateListenerMock) OnICEConnectionStateChange(p *Peer, state webrtc.ICEConnectionState) {
	s.onICEConnectionStateChange(p, state)
}

func (s *StateListenerMock) OnDataChannel(p *Peer, channel *webrtc.DataChannel) {
	s.onDataChannel(p, channel)
}

func TestWebRTC(t *testing.T) {

	ctx := context.Background()
	offererConnectedCh := make(chan bool, 1)
	answererConnectedCh := make(chan bool, 1)

	offererCloseCh := make(chan struct{}, 1)
	answererCloseCh := make(chan struct{}, 1)

	offererDescriptionCh := make(chan string)

	offererControlCh := make(chan struct{})
	answererControlCh := make(chan struct{})

	offererListener := WebrtcListenerMock{}
	answererListener := WebrtcListenerMock{}

	offererStateListener := StateListenerMock{}
	answererStateListener := StateListenerMock{}

	offerer := NewTunnel(ctx, "offerer", []string{"stun:stun.l.google.com:19302"}, &offererListener)
	answerer := NewTunnel(ctx, "answerer", []string{"stun:stun.l.google.com:19302"}, &answererListener)

	var offererDescription string

	offererListener.onLocalDescription = func(p *Peer, sdp webrtc.SessionDescription) {
		slog.Info("offererListener.onLocalDescription", "sdp", sdp.SDP)
		offererDescription = sdp.SDP
	}

	offererListener.onIceCandidate = func(p *Peer, candidate *webrtc.ICECandidate) {
		slog.Info("offererListener.onIceCandidate", "candidate", candidate.ToJSON().Candidate)
		err := answerer.AddICECandidate(candidate.ToJSON().Candidate)
		if err != nil {
			t.Errorf("answerer.AddICECandidate() failed: %v", err)
		}
	}

	answererListener.onLocalDescription = func(p *Peer, sdp webrtc.SessionDescription) {
		slog.Info("answererListener.onLocalDescription", "sdp", sdp.SDP)
		err := offerer.SetRemoteDescription(sdp.SDP)
		close(offererDescriptionCh)
		if err != nil {
			t.Errorf("offerer.SetRemoteDescription() failed: %v", err)
		}
	}

	answererListener.onIceCandidate = func(p *Peer, candidate *webrtc.ICECandidate) {
		<-offererDescriptionCh
		slog.Info("answererListener.onIceCandidate", "candidate", candidate.ToJSON().Candidate)
		err := offerer.AddICECandidate(candidate.ToJSON().Candidate)
		if err != nil {
			t.Errorf("offerer.AddICECandidate() failed: %v", err)
		}
	}

	offererStateListener.onICEConnectionStateChange = func(p *Peer, state webrtc.ICEConnectionState) {
		switch state {
		case webrtc.ICEConnectionStateConnected:
			offererConnectedCh <- true
		case webrtc.ICEConnectionStateFailed:
			offererConnectedCh <- false
		case webrtc.ICEConnectionStateClosed:
			offererCloseCh <- struct{}{}
		}
	}

	offererStateListener.onDataChannel = func(p *Peer, channel *webrtc.DataChannel) {
		if channel.Label() == mainDataChannelLabel {
			offererControlCh <- struct{}{}
		}
	}

	answererStateListener.onICEConnectionStateChange = func(p *Peer, state webrtc.ICEConnectionState) {
		switch state {
		case webrtc.ICEConnectionStateConnected:
			answererConnectedCh <- true
		case webrtc.ICEConnectionStateFailed:
			answererConnectedCh <- false
		case webrtc.ICEConnectionStateClosed:
			answererCloseCh <- struct{}{}
		}
	}

	answererStateListener.onDataChannel = func(p *Peer, channel *webrtc.DataChannel) {
		if channel.Label() == mainDataChannelLabel {
			answererControlCh <- struct{}{}
		}
	}

	offerer.SetStateListener(&offererStateListener)
	answerer.SetStateListener(&answererStateListener)

	err := offerer.Start("")
	if err != nil {
		t.Fatalf("offerer.Start() failed: %v", err)
	}

	if offererDescription == "" {
		t.Fatalf("offererDescription is empty")
	}

	err = answerer.Start(offererDescription)
	if err != nil {
		t.Fatalf("answerer.Start() failed: %v", err)
	}

	if <-offererConnectedCh == false {
		t.Error("offerer failed to connect")
	}

	if <-answererConnectedCh == false {
		t.Error("answerer failed to connect")
	}

	offererControl := false
	answererControl := false
	controlTimeout := time.After(5 * time.Second)

controlLoop:
	for {
		select {
		case <-offererControlCh:
			offererControl = true
			if offererControl && answererControl {
				break controlLoop
			}
		case <-answererControlCh:
			answererControl = true
			if offererControl && answererControl {
				break controlLoop
			}
		case <-controlTimeout:
			t.Errorf("timed out waiting for offerer control channel, offererControl: %t, answererControl: %t", offererControl, answererControl)
			break controlLoop
		}

	}

	if err = offerer.Close(); err != nil {
		t.Errorf("offerer.Close() failed: %v", err)
	}

	if err = answerer.Close(); err != nil {
		t.Errorf("answerer.Close() failed: %v", err)
	}

	timeout := time.After(5 * time.Second)
	offererClosed := false
	answererClosed := false

	for {
		select {
		case <-timeout:
			t.Errorf("timed out waiting for offerer and answerer to close, answererClosed: %t, offererClosed: %t", answererClosed, offererClosed)
			return
		case <-offererCloseCh:
			offererClosed = true
			if offererClosed && answererClosed {
				return
			}
		case <-answererCloseCh:
			answererClosed = true
			if offererClosed && answererClosed {
				return
			}
		}
	}

}
