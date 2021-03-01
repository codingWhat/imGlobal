package main

import (
	"fmt"
	 "github.com/codingWhat/imGlobal/api/config"
	"github.com/codingWhat/imGlobal/api/routers"
	"github.com/codingWhat/imGlobal/common"
	"net/http"
)

func main() {

	//初始化配置文件
	config.InitConfig()

	//初始化redis
	common.InitRedis()

	//初始化Mq
	common.InitMq()

	//启动http服务
	router := routers.New()
	fmt.Println("start 【api server】, addr:", config.G_Config.HttpAddr, "...")
	err := http.ListenAndServe(config.G_Config.HttpAddr, router)
	if err != nil {
		panic(err)
	}
}
