package ws

import (
	"github.com/codingWhat/imGlobal/gateway/config"
	"github.com/gorilla/websocket"
	"net/http"
)

func StartWebSocketServer() {
	WebsocketInit()

	StartClientManager()

	http.HandleFunc("/acc", wsHandler)
	_ = http.ListenAndServe(config.G_Config.WsAddr, nil)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  config.G_Config.WsReadBuffSize,
		WriteBufferSize: config.G_Config.WsWriteBuffSize,
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
