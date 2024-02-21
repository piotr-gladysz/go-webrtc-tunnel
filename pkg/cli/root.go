package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var host string

func CreateCLICommand() *cobra.Command {
	root := &cobra.Command{}

	control := &cobra.Command{
		Use:   "control",
		Long:  "Control commands",
		Short: "Control commands",
	}

	control.AddCommand(controlConnectCmd())
	control.AddCommand(controlDisconnectCmd())

	peer := &cobra.Command{
		Use:   "peer",
		Long:  "Peer commands",
		Short: "Peer commands",
	}

	peer.AddCommand(setPeerCmd())
	peer.AddCommand(removePeerCmd())

	tunnel := &cobra.Command{
		Use:   "tunnel",
		Long:  "Tunnel commands",
		Short: "Tunnel commands",
	}

	tunnel.AddCommand(createTunnelCmd())

	root.AddCommand(control)
	root.AddCommand(peer)
	root.AddCommand(tunnel)

	root.PersistentFlags().StringVar(&host, "host", "127.0.0.1:13080", "host to connect to")
	return root
}

func getHost() string {
	if host != "" {
		return host
	} else {
		return "127.0.0.1:13080"
	}
}

func getHostConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(getHost(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Printf("Failed to connect to host: %v\n", err)
		return nil, err
	}

	return conn, nil
}
