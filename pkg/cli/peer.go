package cli

import (
	"errors"
	"fmt"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/cliapi"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

func setPeerCmd() *cobra.Command {

	var id, portsStr string
	var connect bool

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set peer",
		Long:  "Set peer",
		Run: func(cmd *cobra.Command, args []string) {
			ports, err := parsePorts(portsStr)
			if err != nil {
				fmt.Println("Failed to parse ports: ", err)
				return
			}

			conn, err := getHostConn()
			if err != nil {
				return
			}

			client := cliapi.NewPeerClient(conn)

			req := cliapi.SetPeerRequest{
				Id:      id,
				Ports:   ports,
				Connect: connect,
			}

			_, err = client.SetPeer(cmd.Context(), &req)

			if err != nil {
				fmt.Println("Failed to set peer: ", err)
			} else {
				fmt.Println("Peer set")
			}
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "peer id")
	cmd.Flags().StringVarP(&portsStr, "ports", "p", "", "ports")
	cmd.Flags().BoolVarP(&connect, "connect", "c", false, "connect")

	cmd.MarkFlagRequired("id")

	return cmd
}

func removePeerCmd() *cobra.Command {

	var id string

	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove peer",
		Long:  "Remove peer",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := getHostConn()
			if err != nil {
				return
			}

			client := cliapi.NewPeerClient(conn)

			req := cliapi.RemovePeerRequest{
				Id: id,
			}

			_, err = client.RemovePeer(cmd.Context(), &req)

			if err != nil {
				fmt.Println("Failed to remove peer: ", err)
			} else {
				fmt.Println("Peer removed")
			}
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "peer id")
	cmd.MarkFlagRequired("id")

	return cmd
}

func parsePorts(ports string) ([]uint32, error) {
	ret := make([]uint32, 0)

	commaComponents := strings.Split(ports, ",")

	for _, component := range commaComponents {
		dashComponents := strings.Split(component, "-")
		if len(dashComponents) == 1 {
			port, err := parsePort(dashComponents[0])
			if err != nil {
				return nil, err
			}
			ret = append(ret, port)
		} else if len(dashComponents) == 2 {
			start, err := parsePort(dashComponents[0])
			if err != nil {
				return nil, err
			}
			end, err := parsePort(dashComponents[1])
			if err != nil {
				return nil, err
			}
			for i := start; i <= end; i++ {
				ret = append(ret, i)
			}
		} else {
			return nil, errors.New("invalid port range")
		}
	}

	return ret, nil
}

func parsePort(port string) (uint32, error) {
	val, err := strconv.Atoi(port)

	if err != nil {
		return 0, err
	}

	if val < 0 || val > 65535 {
		return 0, errors.New("invalid port")
	}

	return uint32(val), nil
}
