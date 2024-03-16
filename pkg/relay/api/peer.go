package api

import (
	"context"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/cliapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PeerServer struct {
	cliapi.UnimplementedPeerServer
}

func newPeerServer() *PeerServer {
	return &PeerServer{}
}

func (p *PeerServer) GetPeers(ctx context.Context, req *emptypb.Empty) (*cliapi.PeerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeers not implemented")
}
func (p *PeerServer) SetPeer(ctx context.Context, req *cliapi.SetPeerRequest) (*cliapi.PeerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPeer not implemented")
}
func (p *PeerServer) RemovePeer(ctx context.Context, req *cliapi.RemovePeerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePeer not implemented")
}
