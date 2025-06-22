package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/nitin-kukreti/GoChat/internal/infrastructure/ws"
	"github.com/nitin-kukreti/GoChat/internal/utils"
)

type WebSocketHandler struct {
	Manager *ws.Manager
}

func NewWebSocketHandler(manager *ws.Manager) *WebSocketHandler {
	return &WebSocketHandler{Manager: manager}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true 
	},
}

func (w *WebSocketHandler) HandleConnection(res http.ResponseWriter, req *http.Request) {
	userIDStr := req.URL.Query().Get("userId")
	if userIDStr == "" {
		utils.WriteError(res, http.StatusBadRequest, "userId required")
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		utils.WriteError(res, http.StatusBadRequest, "invalid userId")
		return
	}

	conn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		utils.WriteError(res, http.StatusInternalServerError, "websocket upgrade failed")
		return
	}

	safeConn := ws.NewSafeCon(conn)

	w.Manager.Register(userID, safeConn)

	go w.listen(userID, safeConn)
}


type ChatMessage struct {
	Type utils.MessageType `json:"type"`
	To   int    `json:"to"`  
	Body string `json:"body"`
}


func (w *WebSocketHandler) listen(userID int, conn *ws.SafeConn) {
	defer func() {
		w.Manager.Unregister(userID)
	}()
	w.Manager.Broadcast([]byte(fmt.Sprintf("userId %d join the socket",userID)))

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var parsedMsg ChatMessage;

		if err:=json.Unmarshal(msg,&parsedMsg);err != nil {
			w.Manager.SendMessageToUser(userID,[]byte(fmt.Sprintf("unable to parse message: %s",msg)))
		}

		w.Manager.SendMessageToUser(parsedMsg.To,[]byte(fmt.Sprintf("message from server for msg %s",msg)))
	}
}
