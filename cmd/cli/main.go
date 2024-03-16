package main

import (
	"fmt"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/cli"
)

func main() {
	cmd := cli.CreateCLICommand()

	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
