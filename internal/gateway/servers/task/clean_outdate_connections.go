package task

import (
	"fmt"
	"github.com/codingWhat/imGlobal/internal/gateway/servers/ws"
	"runtime/debug"
)

func cleanConns(params interface{}) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ClearTimeoutConnections stop", r, string(debug.Stack()))
		}
	}()
	ws.ClearOutDateConns()

	return true
}
