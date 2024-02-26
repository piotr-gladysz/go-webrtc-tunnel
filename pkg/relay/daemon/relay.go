package daemon

import (
	"context"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/relay/p2p"
)

type Relay struct {
	ctx context.Context

	signaling *p2p.SignalingClient
}

func NewRelay(ctx context.Context) *Relay {
	return &Relay{
		ctx: ctx,
	}
}
