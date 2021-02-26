package task

import (
	"fmt"
	ws2 "github.com/codingWhat/imGlobal/gateway/servers/ws"
	"runtime/debug"
)

func cleanConns(params interface{}) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ClearTimeoutConnections stop", r, string(debug.Stack()))
		}
	}()
	ws2.ClearOutDateConns()

	return true
}
