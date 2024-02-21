package server

import (
	"context"
	"errors"
	"github.com/gorilla/websocket"
	"go-webrtc-tunnel/pkg/signaling/message"
	"net/http"
	"testing"
	"time"
)

func createSockets(url string, t *testing.T) (*websocket.Conn, *websocket.Conn) {
	sender, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("expected nil, got %s", err)
	}

	receiver, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("expected nil, got %s", err)
	}

	return sender, receiver
}

func receiveToken(conn *websocket.Conn, t *testing.T) string {
	_, data, err := conn.ReadMessage()
	if err != nil {
		t.Fatalf("expected nil, got %s", err)
	}

	var env message.Envelope
	decoder := &message.SimpleJsonDecoder{}
	decoder.Decode(data, &env)

	if env.Type != message.MessageTypeAuthToken {
		t.Fatalf("expected %d, got %d", message.MessageTypeAuthToken, env.Type)
	}

	authMsg, _ := env.GetAuthToken()
	return authMsg.Token
}

func TestSignalingServer(t *testing.T) {

	signaling := NewSignalingServer()

	go func() {
		err := signaling.Start(":18080")
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			t.Fatalf("expected nil, got %s", err)
		}
	}()

	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5)
		err := signaling.Stop(ctx)
		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}
		cancel()

	}()

	u := "ws://127.0.0.1:18080/ws"

	time.Sleep(1 * time.Second)

	sender, receiver := createSockets(u, t)

	defer sender.Close()
	defer receiver.Close()

	senderToken := receiveToken(sender, t)
	receiverToken := receiveToken(receiver, t)

	decoder := &message.SimpleJsonDecoder{}
	encoder := &message.SimpleJsonEncoder{}

	offer := message.NewSDP("offer_string", receiverToken, "", message.MessageTypeSDPOffer)
	data, _ := encoder.Encode(offer)
	err := sender.WriteMessage(websocket.TextMessage, data)

	if err != nil {
		t.Fatalf("expected nil, got %s", err)
	}

	_, data, err = receiver.ReadMessage()
	if err != nil {
		t.Fatalf("expected nil, got %s", err)
	}

	var env message.Envelope
	decoder.Decode(data, &env)

	if env.Type != message.MessageTypeSDPOffer {
		t.Fatalf("expected %d, got %d", message.MessageTypeSDPOffer, env.Type)
	}

	sdp, err := env.GetSDP()

	if err != nil {
		t.Fatalf("expected nil, got %s", err)
	}

	if sdp.SDP != "offer_string" {
		t.Fatalf("expected offer_string, got %s", sdp.SDP)
	}

	if sdp.Sender != senderToken {
		t.Fatalf("expected %s, got %s", senderToken, sdp.Sender)
	}

}
