package daemon

import (
	"context"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/relay/config"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/relay/p2p"
)

type Relay struct {
	cnf *config.Config
	ctx context.Context

	signaling *p2p.SignalingClient
}

func NewRelay(ctx context.Context, cnf *config.Config) *Relay {
	return &Relay{
		cnf: cnf,
		ctx: ctx,
	}
}
