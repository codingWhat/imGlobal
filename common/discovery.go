package common

import (
	"fmt"
	config2 "github.com/codingWhat/imGlobal/internal/gateway/config"
	"runtime/debug"
	"time"
)

const ServiceListKey = "serverList"

func InitDiscovery() {
	InitRedis()
}

func Reg() {
	G_redisClient.HSet(ServiceListKey, config2.G_Config.GrpcAddr, uint64(time.Now().Unix()))
}

func Leave() {
	G_redisClient.HDel(ServiceListKey, config2.G_Config.GrpcAddr)
}

func Discovery(curTime uint64) []string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("service discovery failed, err:", string(debug.Stack()), err)
		}
	}()

	retObj := G_redisClient.HGetAll(ServiceListKey)
	retMap := retObj.Val()

	ret := make([]string, 0)
	for addr := range retMap {
		//todo 过滤心跳超时的server
		ret = append(ret, addr)
	}

	return ret
}
