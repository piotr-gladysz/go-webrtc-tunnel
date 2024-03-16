package api

import (
	"context"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/cliapi"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/relay/daemon"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type RelayControl interface {
}

type ControlServer struct {
	relay *daemon.Relay
	cliapi.UnimplementedControlServer
}

func newControlServer(relay *daemon.Relay) *ControlServer {
	return &ControlServer{
		relay: relay,
	}
}

func (c *ControlServer) Connect(ctx context.Context, req *cliapi.ConnectRequest) (*emptypb.Empty, error) {

	if err := c.relay.StartSignaling(req.Address); err != nil {
		return nil, err
	}

	signaling := c.relay.GetSignaling()

	var err error

	select {
	case <-signaling.WaitForConnectChannel():
	case <-time.After(5 * time.Second):
		err = errors.New("timeout waiting for signaling to connect")
	case <-ctx.Done():
		err = ctx.Err()
	}

	return nil, err
}
func (c *ControlServer) Disconnect(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, c.relay.StopSignaling()
}

func (c *ControlServer) GetStatus(_ context.Context, _ *emptypb.Empty) (*cliapi.RelayStatusResponse, error) {
	ret := &cliapi.RelayStatusResponse{
		Connected:       c.isConnected(),
		RetryCount:      c.retryCount(),
		Uptime:          c.uptime(),
		SignalingServer: c.signalingServer(),
	}

	return ret, nil
}

func (c *ControlServer) isConnected() bool {
	if c.relay.GetSignaling() == nil {
		return false
	}

	return c.relay.GetSignaling().IsConnected()
}

func (c *ControlServer) retryCount() uint32 {
	if c.relay.GetSignaling() == nil {
		return 0
	}

	return c.relay.GetSignaling().GetStatus().GetRetryCount()
}

func (c *ControlServer) uptime() uint32 {
	if c.relay.GetSignaling() == nil || !c.relay.GetSignaling().IsConnected() {
		return 0
	}

	connectTime := c.relay.GetSignaling().GetStatus().GetConnectTime()

	return uint32(time.Since(connectTime).Seconds())
}

func (c *ControlServer) signalingServer() string {
	if c.relay.GetSignaling() == nil {
		return ""
	}

	return c.relay.GetSignaling().GetHost()
}
