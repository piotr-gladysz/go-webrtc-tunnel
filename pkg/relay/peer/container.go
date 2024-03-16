package peer

import "sync"

type PeersContainer struct {
	mux     sync.RWMutex
	allowed map[string]*Peer
}

func NewAllowedPeers() *PeersContainer {
	return &PeersContainer{
		allowed: make(map[string]*Peer),
	}
}

func (a *PeersContainer) Add(id string, peer *Peer) {
	a.mux.Lock()
	defer a.mux.Unlock()

	a.allowed[id] = peer
}

func (a *PeersContainer) Remove(id string) {
	a.mux.Lock()
	defer a.mux.Unlock()

	delete(a.allowed, id)
}

func (a *PeersContainer) Get(id string) *Peer {
	a.mux.RLock()
	defer a.mux.RUnlock()

	return a.allowed[id]
}
