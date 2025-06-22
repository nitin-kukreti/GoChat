package domain

import "github.com/gorilla/websocket"

type WebSocketManager interface {
	Join(userID int, conn *websocket.Conn)
	Leave(userID int)
	SendToUser(userID int, message []byte) error
	Broadcast(message []byte)
}


type WebSocketHub interface {
	Register(userID int, conn *websocket.Conn)
	Unregister(userID int)
	Get(userID int) (*websocket.Conn, bool)
	All() map[int]*websocket.Conn
}
