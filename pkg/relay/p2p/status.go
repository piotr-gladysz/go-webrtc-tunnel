package p2p

import (
	"sync"
	"time"
)

type SignalingStatus struct {
	mux sync.RWMutex

	connectTime           time.Time
	retryCount            int
	isConnected           bool
	lastError             error
	lastConnectionAttempt time.Time
}

func NewSignalingStatus() *SignalingStatus {
	return &SignalingStatus{}
}

func (s *SignalingStatus) SetConnectTime(t time.Time) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.connectTime = t
}

func (s *SignalingStatus) IncRetryCount() {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.retryCount++
}

func (s *SignalingStatus) ResetRetryCount() {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.retryCount = 0
}

func (s *SignalingStatus) SetIsConnected(v bool) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.isConnected = v
}

func (s *SignalingStatus) GetConnectTime() time.Time {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.connectTime
}

func (s *SignalingStatus) GetRetryCount() int {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.retryCount
}

func (s *SignalingStatus) GetIsConnected() bool {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.isConnected
}

func (s *SignalingStatus) SetLastError(err error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.lastError = err
}

func (s *SignalingStatus) GetLastError() error {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.lastError

}

func (s *SignalingStatus) GetLastConnectionAttempt() time.Time {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.lastConnectionAttempt
}

func (s *SignalingStatus) SetLastConnectionAttempt(t time.Time) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.lastConnectionAttempt = t
}
