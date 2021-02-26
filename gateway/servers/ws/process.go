package ws

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/codingWhat/imGlobal/common"
	"github.com/codingWhat/imGlobal/gateway/data/out"
	"sync"
)

var (
	registerHandlers = make(map[string]func(client *Client, seq string, msg []byte) (int, string, interface{}))
	syncLock         = sync.RWMutex{}
)

func Register(cmd string, handler func(client *Client, seq string, msg []byte) (int, string, interface{})) {
	syncLock.Lock()
	defer syncLock.Unlock()
	registerHandlers[cmd] = handler
}

func getHandler(name string) (func(client *Client, seq string, message []byte) (code int, msg string, data interface{}), error) {
	syncLock.RLock()
	defer syncLock.RUnlock()
	rs, ok := registerHandlers[name]
	if !ok {
		return nil, errors.New("handler not exists")
	}

	return rs, nil
}

func ProcessCenter(client *Client, rawData []byte) {
	var reqMsg out.RequestMsg
	err := json.Unmarshal(rawData, &reqMsg)
	if err != nil {
		return
	}

	handler, err := getHandler(reqMsg.Cmd)
	fmt.Println("request Msg:", reqMsg)
	if err != nil {
		return
	}
	requestData, err := json.Marshal(reqMsg.Data)
	if err != nil {
		fmt.Println("处理数据 json Marshal", err)
		return
	}

	code, msg, rs := handler(client, reqMsg.Seq, requestData)
	fmt.Println("handler handle result:", code, msg, rs)
	resp := out.NewResponseDataPack(reqMsg.Seq, reqMsg.Cmd, &common.Response{
		Code: code,
		Msg:  msg,
		Data: rs,
	})
	ret, err := json.Marshal(resp)
	if err != nil {
		return
	}

	client.SendChan <- ret
}
