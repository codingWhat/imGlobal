package ws

import "github.com/gorilla/websocket"

type Bucket struct {
	Conns map[uint64]*websocket.Conn
	Rooms map[uint64]*Room
}

func (b *Bucket) Put() {

}

func (b *Bucket) Get() {

}
