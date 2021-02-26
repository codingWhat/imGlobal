package main

import (
	"fmt"
	"github.com/codingWhat/imGlobal/common"
	"github.com/codingWhat/imGlobal/gateway/config"
	"github.com/codingWhat/imGlobal/gateway/servers/grpcserver"
	"github.com/codingWhat/imGlobal/gateway/servers/task"
	"github.com/codingWhat/imGlobal/gateway/servers/ws"
)

func main() {

	//初始化配置文件
	config.InitConfig()

	//启动redis
	common.InitRedis()

	//启动定时任务
	//启动定时清理超时连接
	//启动服务自动注册和下线
	fmt.Println("start 【task manager】...")
	task.Start()

	//启动grpc服务
	fmt.Println("start 【grpc server】, addr:", config.G_Config.GrpcAddr, "...")
	go grpcserver.StartGrpcServer()

	//启动websocket
	fmt.Println("start 【websocket server】, addr:", config.G_Config.WsAddr, "...")
	ws.StartWebSocketServer()

	//todo
	//拆分 ClientManager-clients
	//Gateway Server 和 api Server 解耦
	//Gateway 增加心跳机制
}
