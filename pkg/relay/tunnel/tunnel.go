package tunnel

import (
	"context"
	"net"
)

type Tunnel struct {
	parentCtx context.Context

	ctx    context.Context
	cancel context.CancelFunc

	id uint16

	port     int
	protocol string
	address  string
	isRemote bool

	conn net.Conn
}

func NewTunnel(parentCtx context.Context, id uint16, port int, protocol, address string, isRemote bool) *Tunnel {
	ctx, cancel := context.WithCancel(parentCtx)
	return &Tunnel{
		parentCtx: parentCtx,
		ctx:       ctx,
		cancel:    cancel,
		id:        id,
		port:      port,
		protocol:  protocol,
		address:   address,
		isRemote:  isRemote,
	}
}

func (t *Tunnel) Start() error {

	panic("implement me")
}
