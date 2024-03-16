package daemon

import (
	"errors"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/relay/p2p"
)

var SignalingAlreadyStarted = errors.New("signaling already started")
var SignalingNotStarted = errors.New("signaling not started")

func (r *Relay) StartSignaling(host string) error {

	if r.signaling != nil {
		return SignalingAlreadyStarted
	}
	r.signaling = p2p.NewSignalingClient(r.ctx, host)
	r.signaling.Start()
	return nil
}

func (r *Relay) StopSignaling() error {

	if r.signaling == nil {
		return SignalingNotStarted
	}

	r.signaling.Stop()
	r.signaling = nil

	return nil
}

func (r *Relay) GetSignaling() *p2p.SignalingClient {
	return r.signaling
}
