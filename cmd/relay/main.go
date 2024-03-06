package main

import (
	"context"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/relay/api"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/relay/config"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/relay/daemon"
	"log/slog"

	"os"
	"os/signal"
)

func main() {

	cnf, err := config.LoadConfig()
	if err != nil {
		slog.Error("failed to load config", "err", err)
		return
	}

	if cnf.Debug {
		slog.SetDefault(createLogger(slog.LevelDebug))
		slog.Debug("Debug mode enabled")
	} else {
		slog.SetDefault(createLogger(slog.LevelInfo))
	}

	ctx := context.Background()
	cancelCtx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	relay := daemon.NewRelay(cancelCtx, cnf)

	server := api.NewServer(cnf.ListenIP, cnf.ListenPort, relay)

	if err := server.Run(); err != nil {
		slog.Error("failed to start server", "err", err)
		return
	}

	relay.StartSignaling(cnf.SignalingHost)

	<-c
	cancel()
	slog.Info("shutting down")
}

func createLogger(level slog.Level) *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: level,
	}

	handler := slog.NewTextHandler(os.Stdout, opts)
	return slog.New(handler)
}
