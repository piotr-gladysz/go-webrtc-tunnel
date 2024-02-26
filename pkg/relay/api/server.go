package api

import (
	"fmt"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/cliapi"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/relay/daemon"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	port int
	ip   string

	relay *daemon.Relay

	server   *grpc.Server
	listener net.Listener
}

func NewServer(port int, ip string, relay *daemon.Relay) *Server {
	return &Server{port: port, ip: ip, relay: relay}
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.ip, s.port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	cliapi.RegisterControlServer(grpcServer, newControlServer())
	cliapi.RegisterPeerServer(grpcServer, newPeerServer())
	cliapi.RegisterTunnelServer(grpcServer, newTunnelServer())

	s.listener = lis
	s.server = grpcServer

	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			return
		}
	}()

	return nil
}
