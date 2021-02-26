package task

import (
	"fmt"
	servers2 "github.com/codingWhat/imGlobal/gateway/servers"
	"runtime/debug"
)

func regServer(params interface{}) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("服务注册 stop", r, string(debug.Stack()))
		}

	}()
	servers2.Reg()
	return true
}

func remServer(params interface{}) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("服务下线 stop", r, string(debug.Stack()))
		}

	}()

	servers2.Leave()
	return true
}
