package daemon

import (
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/relay/p2p"
)

func (r *Relay) StartSignaling(host string) {

	r.signaling = p2p.NewSignalingClient(r.ctx, host)
	r.signaling.Start(r.ctx)
}

func (r *Relay) StopSignaling() {
	r.signaling.Stop()
	r.signaling = nil
}
