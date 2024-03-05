package tunnel

import (
	"context"
	"errors"
	"github.com/pion/webrtc/v4"
	"sync"
)

var AlreadyStartedError = errors.New("tunnel already started")

type TunnelListener interface {
	OnICECandidate(t *Tunnel, candidate *webrtc.ICECandidate)
}

type Tunnel struct {
	id string

	parentCtx context.Context

	ctx    context.Context
	cancel context.CancelFunc

	stunServers []string

	isStarted bool

	lastError   error
	isConnected bool
	connCh      chan struct{}

	peerConnection *webrtc.PeerConnection

	iceCandidates []string
	candidateMux  sync.Mutex
	listener      TunnelListener
}

func NewTunnel(parentCtx context.Context, id string, stunServers []string, listener TunnelListener) *Tunnel {
	return &Tunnel{
		id:          id,
		parentCtx:   parentCtx,
		stunServers: stunServers,
		listener:    listener,
	}
}

func (t *Tunnel) Start(descrStr string) error {
	if t.isStarted {
		return AlreadyStartedError
	}

	t.ctx, t.cancel = context.WithCancel(t.parentCtx)
	t.isStarted = true

	peerConnection, err := webrtc.NewPeerConnection(t.webrtcConfig())
	if err != nil {
		return err
	}

	peerConnection.OnICECandidate(t.onICECandidate)
	peerConnection.OnICEConnectionStateChange(t.onICEConnectionStateChange)

	if descrStr != "" {
		descr := webrtc.SessionDescription{
			Type: webrtc.SDPTypeOffer,
			SDP:  descrStr,
		}
		if err := peerConnection.SetRemoteDescription(descr); err != nil {
			return err
		}
	}
	t.peerConnection = peerConnection

	return nil
}

func (t *Tunnel) webrtcConfig() webrtc.Configuration {
	return webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: t.stunServers,
			},
		},
	}
}

func (t *Tunnel) Close() {

}
