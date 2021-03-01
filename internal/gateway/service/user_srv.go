package service

import (
	"github.com/codingWhat/imGlobal/common"
	"github.com/codingWhat/imGlobal/internal/gateway/config"
	 "github.com/codingWhat/imGlobal/internal/gateway/data/out"
	"strconv"
	"time"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) Login(userInfo out.UserInfo) {
	//保存用户信息uid-userName
	_, _ = us.saveUserInfo(userInfo)
	//存储用户所属grpc服务
	_, _ = us.saveUserServerMap(userInfo)
	//加入聊天室
	_, _ = us.joinRoom(userInfo)
}

func (us *UserService) saveUserInfo(userInfo out.UserInfo) (ret string, err error) {
	key := common.UserInfoRedisPrefixKey + userInfo.UserID
	set := common.G_redisClient.Set(key, userInfo.UserName, 0*time.Second)
	return set.Val(), set.Err()
}

func (us *UserService) joinRoom(userInfo out.UserInfo) (ret bool, err error) {
	key := common.RoomUserListRedisPrefixKey + strconv.Itoa(userInfo.AppID)
	hSet := common.G_redisClient.HSet(key, userInfo.UserID, time.Now().Format("2006-01-02 15:04:05"))

	return hSet.Val(), hSet.Err()
}

func (us *UserService) LeaveRoom(appId,  prefix string) (ret int64, err error) {
	key := common.RoomUserListRedisPrefixKey + appId
	hDel := common.G_redisClient.HDel(key, prefix)
	return hDel.Val(), hDel.Err()
}

func (us *UserService) saveUserServerMap(userInfo out.UserInfo) (ret string, err error) {
	key := common.UserServerMapRedisPrefixKey + strconv.Itoa(userInfo.AppID)
	set := common.G_redisClient.Set(key, config.G_Config.GrpcAddr, 0*time.Second)
	return set.Val(), set.Err()
}
