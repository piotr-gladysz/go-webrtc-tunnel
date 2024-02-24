package api

import (
	"context"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/cliapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ControlServer struct {
	cliapi.UnimplementedControlServer
}

func newControlServer() *ControlServer {
	return &ControlServer{}
}

func (c *ControlServer) Connect(ctx context.Context, req *cliapi.ConnectRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (c *ControlServer) Disconnect(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (c *ControlServer) GetStatus(ctx context.Context, req *emptypb.Empty) (*cliapi.GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
