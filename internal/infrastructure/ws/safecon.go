package ws

import (
	"sync"

	"github.com/gorilla/websocket"
)

type SafeConn struct {
	conn *websocket.Conn
	mu   sync.Mutex
}


func (s *SafeConn) WriteMessage(messageType int, data []byte) error {
	s.mu.Lock();
	defer s.mu.Unlock();
	return s.conn.WriteMessage(messageType,data);

}

func (s *SafeConn) ReadMessage() (messageType int, p []byte, err error) {
 return	s.conn.ReadMessage()
}

func NewSafeCon(conn *websocket.Conn) *SafeConn {
	return &SafeConn{conn: conn}
}