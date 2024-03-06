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
	parentCtx context.Context

	ctx    context.Context
	cancel context.CancelFunc
	host   string

	mux sync.RWMutex

	connMux sync.Mutex
	conn    *websocket.Conn
	connCh  chan struct{}

	decoder server.MessageDecoder
	encoder server.MessageEncoder

	status *SignalingStatus

	reconnectTime time.Duration

	authInfo *message.AuthInfo

	logger *slog.Logger
}

func NewSignalingClient(parentCtx context.Context, host string) *SignalingClient {
	ctx, cancel := context.WithCancel(parentCtx)

	return &SignalingClient{
		parentCtx:     ctx,
		cancel:        cancel,
		host:          host,
		reconnectTime: 2 * time.Second, // TODO: config
		status:        NewSignalingStatus(),
		logger:        slog.Default().With(slog.String("signaling_host", host)),
	}
}

func (s *SignalingClient) Start() {

	s.logger.Info("starting signaling client")

	s.ctx, s.cancel = context.WithCancel(s.parentCtx)
	s.connCh = make(chan struct{})

	go s.watchClose()
	go s.runConnection()

}

func (s *SignalingClient) Stop() {
	s.logger.Info("stopping signaling client")
	s.cancel()
}

func (s *SignalingClient) clean() {
	s.status.SetIsConnected(false)
	s.mux.Lock()
	defer s.mux.Unlock()

	if s.conn != nil {
		s.conn.Close()
	}
	s.conn = nil

}

func (s *SignalingClient) IsConnected() bool {
	return s.status.GetIsConnected()
}

func (s *SignalingClient) WaitForConnectChannel() chan struct{} {
	return s.connCh
}

func (s *SignalingClient) watchClose() {
	<-s.ctx.Done()
	s.clean()
}

func (s *SignalingClient) runConnection() {
	defer func() {
		if s.connCh != nil {
			s.status.SetLastError(ClientStoppedError)
			s.mux.Lock()
			close(s.connCh)
			s.mux.Unlock()

		} else {
			s.logger.Warn("connCh is nil")
		}
	}()

	for {
		s.status.IncRetryCount()
		// prevent reconnecting too fast
		sleepTime := 1*time.Second - time.Since(s.status.GetLastConnectionAttempt())
		if sleepTime > 0 {
			ticker := time.NewTicker(sleepTime)
			select {
			case <-s.ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
			}
		}

		// check if we should stop
		select {
		case <-s.ctx.Done():
			return
		default:
		}

		s.status.SetLastConnectionAttempt(time.Now())

		s.logger.Info("connecting to signaling server")

		conn, err := s.connect()

		// check if we should stop
		select {
		case <-s.ctx.Done():
			return
		default:
		}

		// if we have an error, we should try to reconnect
		if err != nil {
			s.status.SetLastError(err)
			s.logger.Warn("failed to connect to signaling server", "error", err.Error())
			continue
		} else {
			s.status.SetLastError(nil)
			s.status.SetIsConnected(true)
			s.status.SetConnectTime(time.Now())

			s.mux.Lock()
			s.conn = conn
			s.mux.Unlock()

			s.logger.Info("connected to signaling server")
		}

		// if we have a connection, we should process it
		tokenMsg, err := s.recvToken(conn)

		// if we received an error, we should try to reconnect
		if err != nil {
			s.logger.Error("failed to receive token", "error", err.Error())
			continue
		} else {
			s.mux.Lock()
			s.authInfo, _ = tokenMsg.GetAuthToken()
			s.mux.Unlock()
			close(s.connCh)

			s.logger.Info("received token", "token", s.authInfo.Token)
		}

		// if we received a token, we are connected

		s.processWs()

		// if we are not connected, we should try to reconnect
		s.status.SetIsConnected(false)
		s.connCh = make(chan struct{})

		s.mux.Lock()
		s.conn = nil
		s.mux.Unlock()

		// check if we should stop
		select {
		case <-s.ctx.Done():
			return
		default:
		}

		s.logger.Info("disconnected from signaling server, retrying")

	}
}
