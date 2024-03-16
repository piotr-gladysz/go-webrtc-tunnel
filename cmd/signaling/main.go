package main

import (
	"flag"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/signaling/server"
	"log/slog"
	"os"
)

func main() {

	debug := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()

	if *debug {
		slog.SetDefault(createLogger(slog.LevelDebug))
		slog.Debug("Debug mode enabled")
	} else {
		slog.SetDefault(createLogger(slog.LevelInfo))
	}
	//TODO: config

	signaling := server.NewSignalingServer()

	err := signaling.Start(":38080")
	if err != nil {
		panic(err)
	}
}

func createLogger(level slog.Level) *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: level,
	}

	handler := slog.NewTextHandler(os.Stdout, opts)
	return slog.New(handler)
}
