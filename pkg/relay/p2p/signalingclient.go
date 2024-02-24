package p2p

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/signaling/server"
	"sync"
	"time"
)

type SignalingClient struct {
	ctx    context.Context
	cancel context.CancelFunc
	host   string

	mux         sync.RWMutex
	isConnected bool
	conn        *websocket.Conn

	decoder server.MessageDecoder
	encoder server.MessageEncoder

	lastError             error
	lastConnectionAttempt time.Time

	reconnectTime time.Duration
}

func NewSignalingClient(parentCtx context.Context, host string) *SignalingClient {
	ctx, cancel := context.WithCancel(parentCtx)

	return &SignalingClient{
		ctx:           ctx,
		cancel:        cancel,
		host:          host,
		reconnectTime: 2 * time.Second, // TODO: config
	}
}

func (s *SignalingClient) Start(ctx context.Context) error {
	s.ctx = ctx

	go func() {
		select {
		case <-s.ctx.Done():
			s.clean()
			return
		}
	}()

	ch := make(chan struct{})
	go func() {
		s.lastConnectionAttempt = time.Now()
		conn, err := s.connect()

		// if first connection failed, don't try to reconnect

		s.mux.Lock()
		if err != nil {
			s.lastError = err
			ch <- struct{}{}
			s.mux.Unlock()
			return
		} else {
			s.lastError = nil
			ch <- struct{}{}
		}

		s.isConnected = true
		s.conn = conn
		s.mux.Unlock()

		for {

			select {
			case <-s.ctx.Done():
				return
			default:
			}

			s.processWs(conn)

			s.mux.Lock()
			s.isConnected = false
			s.conn = nil
			s.mux.Unlock()

			for {

				sleepTime := 1*time.Second - time.Since(s.lastConnectionAttempt)
				if sleepTime > 0 {
					ticker := time.NewTicker(sleepTime)
					select {
					case <-s.ctx.Done():
						ticker.Stop()
						return
					case <-ticker.C:
					}
				}

				select {
				case <-s.ctx.Done():
					return
				default:
				}

				s.lastConnectionAttempt = time.Now()

				conn, err := s.connect()

				select {
				case <-s.ctx.Done():
					return
				default:
				}

				s.mux.Lock()
				if err != nil {
					s.lastError = err
				} else {
					s.lastError = nil
					s.isConnected = true
					s.conn = conn
					s.mux.Unlock()
					break
				}
				s.mux.Unlock()

			}

		}

	}()

	<-ch

	return s.lastError
}

func (s *SignalingClient) Stop() {
	s.cancel()
}

func (s *SignalingClient) clean() {
	s.mux.Lock()
	defer s.mux.Unlock()

	if s.conn != nil {
		s.conn.Close()
	}
	s.conn = nil

	s.isConnected = false
}

func (s *SignalingClient) IsConnected() bool {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.isConnected
}
