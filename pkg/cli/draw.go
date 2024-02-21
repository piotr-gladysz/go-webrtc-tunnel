package cli

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/cliapi"
	"io"
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
