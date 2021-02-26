package common

import (
	"github.com/go-redis/redis"
)

var G_redisClient *redis.Client

func InitRedis() {
	G_redisClient = redis.NewClient(&redis.Options{
		Addr:         ":6379",
		Password:     "",
		PoolSize:     20,
		MinIdleConns: 10,
	})
}

const (
	UserInfoRedisPrefixKey      = "USERINFO:"
	RoomUserListRedisPrefixKey  = "ROOMUSERLIST:"
	UserServerMapRedisPrefixKey = "USERSERVERMAP:"
)
