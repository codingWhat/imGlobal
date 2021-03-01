package models

import (
	"github.com/codingWhat/imGlobal/api/data"
	"github.com/codingWhat/imGlobal/common"
	"time"
)

type UserModel struct {
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (um *UserModel) getRedisKey(userId string) string {
	return "USERINFO:" + userId
}

func (um *UserModel) GetUserInfo(userId string) (ret string, err error) {
	get := common.G_redisClient.Get(um.getRedisKey(userId))
	return get.Val(), get.Err()
}

func GetUserInfo(userId string) (ret string, err error) {
	return NewUserModel().GetUserInfo(userId)
}

func (um *UserModel) GetRoomUserRedisKey(appId string) string {
	return common.RoomUserListRedisPrefixKey + appId
}

func (um *UserModel) GetRoomUsers(appId string) (ret map[string][]data.UserLoginInfo, err error) {
	key := um.GetRoomUserRedisKey(appId)
	rsObj := common.G_redisClient.HGetAll(key)

	ret = make(map[string][]data.UserLoginInfo)
	ret["userList"] = make([]data.UserLoginInfo, 0)
	_, _ = time.LoadLocation("asia/shanghai")
	curr := time.Now()
	for userId, lastTime := range rsObj.Val() {
		loc, _ := time.LoadLocation("Local")
		st, _ := time.ParseInLocation("2006-01-02 15:04:05", lastTime, loc)
		userName, _ := um.GetUserInfo(userId)
		loginInfo := data.UserLoginInfo{
			UserId:userId,
			UserName: userName,
		}

		if curr.Sub(st).Minutes() < 1  {
			loginInfo.IsAlive = true
		}
		ret["userList"] = append(ret["userList"], loginInfo)

	}

	return ret, rsObj.Err()
}
