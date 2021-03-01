package ws

import (
	"fmt"
	"github.com/codingWhat/imGlobal/common"
	"sync"
	"time"
)

var G_clientManager *ClientManager

func StartClientManager() {
	G_clientManager = NewClientManager()
	go G_clientManager.Run()
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		Clients:   make(map[*Client]bool),
		regChan:   make(chan *Client, 100),
		broadChan: make(chan []byte, 100),
		unRegChan: make(chan *Client, 100),
	}
}

type ClientManager struct {
	//todo 优化点 拆分client,  降低锁粒度
	Clients map[*Client]bool
	sync.RWMutex
	regChan   chan *Client
	unRegChan chan *Client
	broadChan chan []byte
}

func (cm *ClientManager) Run() {

	for {
		select {
		case client := <-cm.regChan:
			cm.regClient(client)

		case msg := <-cm.broadChan:
			fmt.Println("broad--->")
			cm.broadMsg(msg)

		case client := <-cm.unRegChan:
			fmt.Println("unregister user:", client.UserId)
			delete(cm.Clients, client)
			client.Close()
			common.G_redisClient.HDel("USERLIST:101", client.UserId)
		}
	}
}

func (cm *ClientManager) GetCurrentClients() map[*Client]bool {
	cm.RLock()
	defer cm.RUnlock()
	return cm.Clients
}

func (cm *ClientManager) regClient(client *Client) {
	cm.Lock()
	defer cm.Unlock()
	cm.Clients[client] = true

	fmt.Println("current UserMap, ", cm.Clients)
}

func (cm *ClientManager) broadMsg(msg []byte) {
	cm.RLock()
	defer cm.RUnlock()
	for client := range cm.Clients {
		client.SendChan <- msg
	}
}

func ClearOutDateConns() {
	curTime := uint64(time.Now().UnixNano())
	for client := range G_clientManager.Clients {
		if (client.LastHeartBeat + 6*60) <= curTime {
			G_clientManager.unRegChan <- client
		}
	}
}
