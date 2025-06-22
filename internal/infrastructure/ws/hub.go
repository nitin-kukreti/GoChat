package ws

import (
	"sync"
)

type Hub struct {
	clients map[int]*SafeConn
	lock    sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[int]*SafeConn),
	}
}

func (h *Hub) RegisterClient(userID int, conn *SafeConn) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.clients[userID] = conn
}

func (h *Hub) GetClient(userID int) (*SafeConn, bool) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	conn, exists := h.clients[userID]
	return conn, exists
}

func (h *Hub) RemoveClient(userID int) {
	h.lock.Lock()
	defer h.lock.Unlock()
	delete(h.clients, userID)
}


func (h *Hub) AllClients() map[int]*SafeConn {
	h.lock.RLock()
	defer h.lock.RUnlock()
	copy := make(map[int]*SafeConn, len(h.clients))
	for k, v := range h.clients {
		copy[k] = v
	}
	return copy
}