package tunnel

import (
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/relay/proxymsg"
	"sync"
)

type Controller struct {
	writeMux sync.Mutex
	readMux  sync.Mutex

	tunnelsMap map[uint16]*Tunnel

	WriteMessageHandler func(msg proxymsg.ProxyMessage) error
}

func NewController() *Controller {
	return &Controller{
		tunnelsMap: make(map[uint16]*Tunnel),
	}
}

func (c *Controller) RecvMessage(msg proxymsg.ProxyMessage) error {
	return nil
}
