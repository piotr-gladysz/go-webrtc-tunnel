package main

import (
	"context"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/relay/api"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/relay/daemon"
	"log/slog"

	"os"
	"os/signal"
)

func main() {
	ctx := context.Background()
	cancelCtx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	relay := daemon.NewRelay(cancelCtx)

	server := api.NewServer(8080, "127.0.0.1", relay)

	if err := server.Run(); err != nil {
		slog.Error("failed to start server", "err", err)
		return
	}

	// TODO: remove hardcoded signaling
	relay.StartSignaling("ws://127.0.0.1:38080/ws")

	<-c
	cancel()
	slog.Info("shutting down")

}
