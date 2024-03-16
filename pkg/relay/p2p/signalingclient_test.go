package p2p

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/signaling/server"
	"io"
	"log/slog"
	"net/http"
	"testing"
	"time"
)

func init() {
	fmt.Printf("SignalingClient tests init, slog is disabled\n")
	slog.SetDefault(slog.New(slog.NewTextHandler(io.Discard, nil)))

	gin.SetMode(gin.TestMode)
}

func TestSignalingClient_Start(t *testing.T) {
	srvCh := make(chan struct{})
	var srv *server.SignalingServer

	go func() {
		defer close(srvCh)
		srv = server.NewSignalingServer()
		err := srv.Start(":18081")
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			t.Error(err)
		} else {
			t.Log("Server closed")
		}
	}()

	time.Sleep(1 * time.Second)

	ctx := context.Background()
	client := NewSignalingClient(ctx, "localhost:18081")

	client.Start()

	select {
	case <-client.connCh:
		t.Log("Connected")
	case <-time.After(10 * time.Second):
		t.Fatal("Timeout")
	}

	if client.authInfo == nil {
		t.Error("AuthInfo is nil")
	} else {
		if client.authInfo.Token == "" {
			t.Error("Token is empty")
		} else {
			t.Log("Token:", client.authInfo.Token)
		}
	}

	if client.status.GetLastConnectionAttempt().IsZero() {
		t.Error("LastConnectionAttempt is 0")
	}

	if client.status.GetConnectTime().IsZero() {
		t.Error("ConnectTime is 0")
	}

	if client.status.GetLastError() != nil {
		t.Error("Error is not nil")
	}

	if client.status.GetRetryCount() != 1 {
		t.Error("RetryCount is not 1")
	}

	client.Stop()
	srv.Stop(ctx)

	<-srvCh
}
