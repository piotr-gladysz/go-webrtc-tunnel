package api

import (
	"context"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/cliapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TunnelServer struct {
	cliapi.UnimplementedTunnelServer
}

func newTunnelServer() *TunnelServer {
	return &TunnelServer{}
}

func (t *TunnelServer) Create(ctx context.Context, req *cliapi.CreateTunnelRequest) (*cliapi.TunnelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (t *TunnelServer) List(ctx context.Context, req *emptypb.Empty) (*cliapi.ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (t *TunnelServer) Delete(ctx context.Context, req *cliapi.DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
