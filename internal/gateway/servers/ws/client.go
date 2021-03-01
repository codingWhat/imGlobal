package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn          *websocket.Conn
	cm            *ClientManager
	SendChan      chan []byte
	closeChan     chan bool
	AppId         int
	UserId        string
	LastHeartBeat uint64
	isClosed      bool
}

func NewClient(cm *ClientManager, conn *websocket.Conn) *Client {
	return &Client{
		cm:        cm,
		conn:      conn,
		SendChan:  make(chan []byte, 100),
		closeChan: make(chan bool),
	}
}

func (c *Client) Read() {

	for {
		_, data, err := c.conn.ReadMessage()
		if err != nil {
			fmt.Println("read from ws failed. err:", err.Error())
			c.cm.unRegChan <- c
			fmt.Println("has sent to unRegChan")
			//c.Close()
			return
		}

		//c.SendChan <- data
		fmt.Println("read data:", string(data))
		ProcessCenter(c, data)
	}
}

func (c *Client) Write() {

	for {
		select {
		case data := <-c.SendChan:
			_ = c.conn.WriteMessage(websocket.TextMessage, data)
		}
	}
}

func (c *Client) Close() {
	if c.isClosed {
		return
	}

	close(c.closeChan)
	close(c.SendChan)
	_ = c.conn.Close()

	c.isClosed = true
}

