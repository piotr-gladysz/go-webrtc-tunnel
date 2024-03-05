package api

import (
	"context"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/cliapi"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/relay/daemon"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
func (c *ControlServer) Disconnect(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, c.relay.StopSignaling()
}
func (c *ControlServer) GetStatus(ctx context.Context, req *emptypb.Empty) (*cliapi.GetStatusResponse, error) {
	//ret := cliapi.GetStatusResponse{
	//	Status:     "",
	//	RetryCount: "",
	//	Uptime:     "",
	//}

	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
