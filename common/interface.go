package common

type PushMsg struct {
	Destination string
	Value []byte
}

type ConsumeMsg struct {
	Source interface{}
	Data interface{}
}

type MQ interface {
	Push (msg PushMsg)
	Pull (msg ConsumeMsg)
}