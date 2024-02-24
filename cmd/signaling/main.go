package main

import (
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/signaling/server"
)

func main() {

	//TODO: config
	signaling := server.NewSignalingServer()

	err := signaling.Start(":38080")
	if err != nil {
		panic(err)
	}
}
