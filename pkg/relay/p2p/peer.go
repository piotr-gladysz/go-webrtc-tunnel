package p2p

import "sync"

type AllowedPeers struct {
	mux     sync.RWMutex
	allowed map[string][]int
}

func NewAllowedPeers() *AllowedPeers {
	return &AllowedPeers{
		allowed: make(map[string][]int),
	}
}

func (a *AllowedPeers) Add(peer string, ports []int) {
	a.mux.Lock()
	defer a.mux.Unlock()

	a.allowed[peer] = ports
}

func (a *AllowedPeers) Remove(peer string) {
	a.mux.Lock()
	defer a.mux.Unlock()

	delete(a.allowed, peer)
}

func (a *AllowedPeers) Get(peer string) []int {
	a.mux.RLock()
	defer a.mux.RUnlock()

	return a.allowed[peer]
}

func (a *AllowedPeers) IsAllowed(peer string, port int) bool {
	ports := a.Get(peer)
	for _, p := range ports {
		if p == port {
			return true
		}
	}
	return false
}
