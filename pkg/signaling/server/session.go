package server

import (
	"github.com/gorilla/websocket"
	"log/slog"
	"sync"
)

type WSSession struct {
	id   string
	conn *websocket.Conn
}

type SessionStorage struct {
	mux           *sync.RWMutex
	connToSession map[*websocket.Conn]*WSSession

	idToSession map[string]*WSSession
}

func NewSessionStorage() *SessionStorage {
	return &SessionStorage{
		mux:           &sync.RWMutex{},
		connToSession: make(map[*websocket.Conn]*WSSession),
		idToSession:   make(map[string]*WSSession),
	}
}

func (s *SessionStorage) Add(session *WSSession) {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.connToSession[session.conn] = session
	s.idToSession[session.id] = session

}

func (s *SessionStorage) Remove(session *WSSession) {
	s.mux.Lock()
	defer s.mux.Unlock()

	delete(s.connToSession, session.conn)
	delete(s.idToSession, session.id)
}

func (s *SessionStorage) GetById(id string) *WSSession {
	s.mux.RLock()
	defer s.mux.RUnlock()

	return s.idToSession[id]
}

func (s *SessionStorage) GetByConn(conn *websocket.Conn) *WSSession {
	s.mux.RLock()
	defer s.mux.RUnlock()

	return s.connToSession[conn]
}

func (w *WSSession) Send(data []byte) bool {
	if err := w.conn.WriteMessage(websocket.TextMessage, data); err != nil {
		slog.Error("Failed to send message", "err", err)
		return false
	}

	return true
}
