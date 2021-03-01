package ws

import (
	config2 "github.com/codingWhat/imGlobal/internal/gateway/config"
	"github.com/gorilla/websocket"
	"net/http"
)

func StartWebSocketServer() {
	WebsocketInit()

	StartClientManager()

	http.HandleFunc("/acc", wsHandler)
	_ = http.ListenAndServe(config2.G_Config.WsAddr, nil)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  config2.G_Config.WsReadBuffSize,
		WriteBufferSize: config2.G_Config.WsWriteBuffSize,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	client := NewClient(G_clientManager, conn)
	G_clientManager.regChan <- client

	go client.Read()
	go client.Write()
}
