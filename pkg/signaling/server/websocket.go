package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/signaling/message"
	"log/slog"
	"time"
)

type MessageEncoder interface {
	Encode(msg *message.Envelope) ([]byte, error)
}

type MessageDecoder interface {
	Decode(data []byte, msg *message.Envelope) error
}

func (s *SignalingServer) handleWS(c *gin.Context) {

	slog.Debug("New WS connection", "remote", c.Request.RemoteAddr)
	session, err := s.authenticate(c)
	if err != nil {
		slog.Warn("Failed to authenticate WS", "err", err)
		c.AbortWithStatus(401)
		return
	}

	conn, err := s.upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		slog.Error("Failed to upgrade WS", "err", err)
		c.AbortWithStatus(500)
		return
	}

	session.conn = conn

	slog.Debug("WS connection upgraded", "session", session.id, "remote", conn.RemoteAddr().String())

	go s.handleWsConn(session)
}

func (s *SignalingServer) authenticate(c *gin.Context) (*WSSession, error) {
	// TODO: implement authentication
	// TODO: public id, private token, maybe asymmetric encryption
	// TODO: sticky tokens

	// !!! Temporary implementation
	sess := &WSSession{
		id: fmt.Sprintf("%d", time.Now().UnixNano()),
	}

	return sess, nil

}

func (s *SignalingServer) handleWsConn(session *WSSession) {
	slog.Debug("New WS connection", "session", session.id, "remote", session.conn.RemoteAddr())

	s.sessions.Add(session)

	defer s.sessions.Remove(session)

	authMsg := message.NewAuthToken(session.id)
	encoded, err := s.encoder.Encode(authMsg)

	if err != nil {
		slog.Error("Failed to encode auth message", "err", err)
		return
	}

	if !session.Send(encoded) {
		return
	}

	for {
		read, rawMsg, err := session.conn.ReadMessage()
		if err != nil {
			slog.Error("Failed to read message", "err", err, "session", session.id, "read", read)
			return
		}
		slog.Debug("received message", "message", string(rawMsg))

		s.handleMessage(session, rawMsg)

	}
}

func (s *SignalingServer) handleMessage(session *WSSession, rawMsg []byte) error {
	var env message.Envelope

	err := s.decoder.Decode(rawMsg, &env)
	if err != nil {
		slog.Error("Failed to decode message", "err", err)

		errMsg := message.NewErrorMessage("Invalid message", 1, -1)
		encodedErr, err := s.encoder.Encode(errMsg)
		if err != nil {
			slog.Error("Failed to encode error message", "err", err)
			return err
		}

		if !session.Send(encodedErr) {
			return err
		}
	}

	switch env.Type {
	case message.MessageTypeSDPOffer, message.MessageTypeSDPAnswer:
		err = s.handleSignaling(session, &env)
		if err != nil {
			slog.Error("Failed to handle Signaling", "err", err)
		}

	}
	return nil
}
