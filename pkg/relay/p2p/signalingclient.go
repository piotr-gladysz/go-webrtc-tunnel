package p2p

import (
	"context"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/signaling/message"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/signaling/server"
	"log/slog"
	"sync"
	"time"
)

var ClientStoppedError = errors.New("client stopped")

type SignalingClient struct {
	ctx    context.Context
	cancel context.CancelFunc
	host   string

	mux         sync.RWMutex
	isConnected bool
	conn        *websocket.Conn
	connCh      chan struct{}

	decoder server.MessageDecoder
	encoder server.MessageEncoder

	lastError             error
	lastConnectionAttempt time.Time

	reconnectTime time.Duration

	authInfo *message.AuthInfo
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

func (s *SignalingClient) Start(ctx context.Context) {
	s.ctx = ctx

	s.connCh = make(chan struct{})

	go func() {
		select {
		case <-s.ctx.Done():
			s.clean()
			return
		}
	}()

	go func() {
		defer func() {
			if s.connCh != nil {
				s.mux.Lock()
				s.lastError = ClientStoppedError
				close(s.connCh)
				s.mux.Unlock()

			} else {
				slog.Warn("connCh is nil")
			}
		}()

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
				s.mux.Unlock()
				continue
			} else {
				s.lastError = nil
				s.isConnected = true
				s.conn = conn
			}
			s.mux.Unlock()

			tokenMsg, err := s.recvToken(conn)
			if err != nil {
				slog.Error("failed to receive token", "error", err.Error())
				continue
			} else {
				s.mux.Lock()
				s.authInfo, _ = tokenMsg.GetAuthToken()
				close(s.connCh)
				s.mux.Unlock()
			}

			s.processWs()

			s.connCh = make(chan struct{})
			s.mux.Lock()
			s.isConnected = false
			s.conn = nil
			s.mux.Unlock()

		}

	}()

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

func (s *SignalingClient) WaitForConnectChannel() chan struct{} {
	return s.connCh
}
