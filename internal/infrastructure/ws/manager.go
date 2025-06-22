package ws

import (
	"github.com/gorilla/websocket"
)

type Manager struct {
	hub *Hub
}

func NewManager() *Manager {
	return &Manager{
		hub: NewHub(),
	}
}

func (m *Manager) Register(userID int, conn *SafeConn) {
	m.hub.RegisterClient(userID, conn)
}

func (m *Manager) Unregister(userID int) {
	m.hub.RemoveClient(userID)
}

func (m *Manager) SendMessageToUser(userID int, msg []byte) error {
	conn, ok := m.hub.GetClient(userID)
	if !ok {
		return nil // or return error
	}
	return conn.WriteMessage(websocket.TextMessage, msg)
}

func (m *Manager) Broadcast(msg []byte) {
	clients := m.hub.AllClients() // optional method to return a copy
	for _, conn := range clients {
		conn.WriteMessage(websocket.TextMessage, msg)
	}
}
