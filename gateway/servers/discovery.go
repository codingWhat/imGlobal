package servers

import (
	"fmt"
	"github.com/codingWhat/imGlobal/common"
	"github.com/codingWhat/imGlobal/gateway/config"
	"runtime/debug"
	"time"
)

const ServiceListKey = "serverList"

func Reg() {
	common.G_redisClient.HSet(ServiceListKey, config.G_Config.GrpcAddr, uint64(time.Now().Unix()))
}

func Leave() {
	common.G_redisClient.HDel(ServiceListKey, config.G_Config.GrpcAddr)
}

func Discovery(curTime uint64) []string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("service discovery failed, err:", string(debug.Stack()))
		}
	}()

	retObj := common.G_redisClient.HGetAll(ServiceListKey)
	retMap := retObj.Val()

	ret := make([]string, 0)
	for addr := range retMap {

		ret = append(ret, addr)
	}

	return ret
}
