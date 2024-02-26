package p2p

import (
	"context"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/signaling/message"
	"log/slog"
)

func (s *SignalingClient) connect() (*websocket.Conn, error) {
	conn, _, err := websocket.DefaultDialer.Dial(s.host, nil)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (s *SignalingClient) recvToken(conn *websocket.Conn) (*message.Envelope, error) {
	_, data, err := conn.ReadMessage()
	if err != nil {
		return nil, err
	}

	var env message.Envelope
	decoder := &message.SimpleJsonDecoder{}
	if err := decoder.Decode(data, &env); err != nil {
		return nil, err
	}

	if env.Type != message.MessageTypeAuthToken {
		return nil, errors.New("expected auth token")
	}
	return &env, nil
}

func (s *SignalingClient) waitForToken(ctx context.Context, conn *websocket.Conn) (*message.Envelope, error) {

	okChan := make(chan *message.Envelope)
	errChan := make(chan error)

	go func() {
		ret, err := s.recvToken(conn)
		if err != nil {
			errChan <- err
		} else {
			okChan <- ret
		}
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case env := <-okChan:
		return env, nil
	case err := <-errChan:
		return nil, err
	}
}

func (s *SignalingClient) processWs() {
	for {

		read, rawMsg, err := s.conn.ReadMessage()
		if err != nil {
			slog.Error("Failed to read message", "err", err, "raw", string(rawMsg), "read", read)
			return
		}
		slog.Debug("received message", "message", string(rawMsg))

		err = s.handleMessage(rawMsg)
		if err != nil {
			slog.Error("Failed to handle message", "err", err)
		}

	}
}

func (s *SignalingClient) handleMessage(rawMsg []byte) error {
	var env message.Envelope

	err := s.decoder.Decode(rawMsg, &env)
	if err != nil {
		slog.Error("Failed to decode message", "err", err)

		if err := s.sendError("Invalid message", 1, -1); err != nil {
			slog.Error("Failed to send error", "err", err)
			return err
		}
	}

	switch env.Type {
	case message.MessageTypeSDPOffer, message.MessageTypeSDPAnswer:

	}
	return nil
}

func (s *SignalingClient) sendMsg(msg *message.Envelope) error {
	encoded, err := s.encoder.Encode(msg)
	if err != nil {
		return err
	}

	err = s.conn.WriteMessage(websocket.TextMessage, encoded)
	if err != nil {
		return err
	}

	return nil
}

func (s *SignalingClient) sendError(errMsg string, code, relatedId int32) error {
	msg := message.NewErrorMessage(errMsg, code, relatedId)
	return s.sendMsg(msg)
}
