package task

import (
	"fmt"
	"github.com/codingWhat/imGlobal/common"
	"runtime/debug"
)

func regServer(params interface{}) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("服务注册 stop", r, string(debug.Stack()))
		}

	}()
	common.Reg()
	return true
}

func remServer(params interface{}) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("服务下线 stop", r, string(debug.Stack()))
		}

	}()

	common.Leave()
	return true
}
