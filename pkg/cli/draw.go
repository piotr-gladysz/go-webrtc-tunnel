package cli

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/cliapi"
	"io"
	"strings"
)

func drawTunnel(writer io.Writer, tunnel *cliapi.TunnelResponse) {
	t := table.NewWriter()
	t.SetOutputMirror(writer)
	t.AppendHeader(table.Row{"Id", "PeerId", "LocalPort", "RemotePort"})

	t.AppendRow(
		table.Row{
			tunnel.Id,
			tunnel.PeerId,
			tunnel.LocalPort,
			tunnel.RemotePort,
		})

	t.Render()
}

func drawStatus(writer io.Writer, status *cliapi.RelayStatusResponse) {
	t := table.NewWriter()
	t.SetOutputMirror(writer)
	t.AppendHeader(table.Row{"Name", "Value"})

	t.AppendRows([]table.Row{
		{
			"Connected", status.Connected,
		},
		{
			"RetryCount", status.RetryCount,
		},
		{
			"Uptime", fmt.Sprintf("%ds", status.Uptime),
		},
		{
			"SignalingServer", status.SignalingServer,
		},
	})

	t.Render()
}

func drawPeer(writer io.Writer, peer *cliapi.PeerResponse) {
	t := table.NewWriter()
	t.SetOutputMirror(writer)

	t.AppendHeader(table.Row{"Name", "Value"})

	t.AppendRows([]table.Row{
		{
			"Id", peer.Id,
		},
		{
			"Connected", peer.Connected,
		},
		{
			"Ports", strings.Join(peer.Ports, ","),
		},
		{
			"Local ports", strings.Join(peer.LocalPorts, ","),
		},
		{
			"Remote ports", strings.Join(peer.RemotePorts, ","),
		},
	})

	t.Render()
}

func drawPeers(writer io.Writer, peers []*cliapi.PeerResponse) {
	t := table.NewWriter()
	t.SetOutputMirror(writer)

	t.AppendHeader(table.Row{"Id", "Connected", "Ports", "Local ports", "Remote ports"})

	for _, peer := range peers {
		t.AppendRow(
			table.Row{
				peer.Id,
				peer.Connected,
				strings.Join(peer.Ports, ","),
				strings.Join(peer.LocalPorts, ","),
				strings.Join(peer.RemotePorts, ","),
			})
	}

	t.Render()
}
