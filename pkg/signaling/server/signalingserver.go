package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/signaling/message"
	"log/slog"
	"net/http"
)

type SignalingServer struct {
	server *http.Server

	upgrader *websocket.Upgrader
	encoder  MessageEncoder
	decoder  MessageDecoder

	sessions *SessionStorage
}

func NewSignalingServer() *SignalingServer {

	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	encoder := &message.SimpleJsonEncoder{}
	decoder := &message.SimpleJsonDecoder{}

	return &SignalingServer{
		upgrader: upgrader,
		encoder:  encoder,
		decoder:  decoder,
		sessions: NewSessionStorage()}
}

func (s *SignalingServer) Start(listenAddr string) error {

	slog.Info("starting signaling server", "listen_addr", listenAddr)

	router := s.createRouter()

	server := &http.Server{
		Addr:    listenAddr,
		Handler: router,
	}

	s.server = server

	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *SignalingServer) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *SignalingServer) createRouter() *gin.Engine {
	ret := gin.Default()

	ret.GET("/ws", s.handleWS)

	return ret
}
