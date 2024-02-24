package p2p

import (
	"context"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/signaling/message"
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

func (s *SignalingClient) processWs(conn *websocket.Conn) {

}
