package main

import (
	"go-webrtc-tunnel/pkg/signaling/server"
)

func main() {

	signaling := server.NewSignalingServer()

	err := signaling.Start(":8080")
	if err != nil {
		panic(err)
	}
}
