package models

import (
	"github.com/codingWhat/imGlobal/common"
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

func (um *UserModel) GetRoomUsers(appId string) (ret map[string][]string, err error) {
	key := um.GetRoomUserRedisKey(appId)
	rsObj := common.G_redisClient.HGetAll(key)

	ret = make(map[string][]string)
	ret["userList"] = make([]string, 0)
	for user := range rsObj.Val() {
		ret["userList"] = append(ret["userList"], user)
	}

	return ret, rsObj.Err()
}
