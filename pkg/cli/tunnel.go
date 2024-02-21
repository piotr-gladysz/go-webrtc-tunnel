package cli

import (
	"fmt"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/cliapi"
	"github.com/spf13/cobra"
)

func createTunnelCmd() *cobra.Command {
	var perId string
	var localPort, remotePort uint32

	ret := &cobra.Command{
		Use:   "create",
		Short: "Create tunnel",
		Long:  "Create tunnel",
		Run: func(cmd *cobra.Command, args []string) {
			host, err := getHostConn()
			if err != nil {
				return
			}

			client := cliapi.NewTunnelClient(host)

			req := &cliapi.CreateTunnelRequest{
				PeerId:     perId,
				LocalPort:  localPort,
				RemotePort: remotePort,
			}

			ret, err := client.Create(cmd.Context(), req)

			if err != nil {
				fmt.Println("Failed to create tunnel: ", err)
				return
			} else {
				drawTunnel(cmd.OutOrStdout(), ret)
			}

		},
	}

	ret.Flags().StringVarP(&perId, "peer-id", "i", "", "peer id")
	ret.Flags().Uint32VarP(&localPort, "local-port", "l", 0, "local port")
	ret.Flags().Uint32VarP(&remotePort, "remote-port", "r", 0, "remote port")

	return ret
}
