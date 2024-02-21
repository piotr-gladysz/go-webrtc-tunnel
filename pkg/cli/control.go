package cli

import (
	"fmt"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/cliapi"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/emptypb"
)

func controlConnectCmd() *cobra.Command {
	var address string

	cmd := &cobra.Command{
		Use:   "connect",
		Long:  "Connect to the control server",
		Short: "Connect to the control server",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := getHostConn()
			if err != nil {
				return
			}

			client := cliapi.NewControlClient(conn)

			req := &cliapi.ConnectRequest{
				Address: address,
			}

			_, err = client.Connect(cmd.Context(), req)

			if err != nil {
				fmt.Println("Failed to connect to control server: ", err)
			} else {
				fmt.Println("Connected to control server")
			}

			defer conn.Close()
		},
	}

	cmd.Flags().StringVarP(&address, "address", "a", "", "signaling server address")
	cmd.MarkFlagRequired("address")

	return cmd
}

func controlDisconnectCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "disconnect",
		Long:  "Disconnect from the control server",
		Short: "Disconnect from the control server",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := getHostConn()
			if err != nil {
				return
			}

			client := cliapi.NewControlClient(conn)

			_, err = client.Disconnect(cmd.Context(), &emptypb.Empty{})

			if err != nil {
				fmt.Println("Failed to disconnect from control server: ", err)
			} else {
				fmt.Println("Disconnected from control server")
			}

			defer conn.Close()
		},
	}
}

func controlGetStatusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Long:  "Get status of the control server",
		Short: "Get status of the control server",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := getHostConn()
			if err != nil {
				return
			}

			client := cliapi.NewControlClient(conn)

			_, err = client.GetStatus(cmd.Context(), &emptypb.Empty{})

			if err != nil {
				fmt.Println("Failed to get status of control server: ", err)
			} else {
				fmt.Println("Control server is running")
			}

			defer conn.Close()
		},
	}
}
