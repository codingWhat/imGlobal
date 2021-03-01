package ws

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/codingWhat/imGlobal/common"
	"github.com/codingWhat/imGlobal/internal/gateway/config"
	"github.com/codingWhat/imGlobal/internal/gateway/data/out"
	"github.com/codingWhat/imGlobal/internal/gateway/service"
	"github.com/codingWhat/imGlobal/protobuf"
	"github.com/golang/protobuf/proto"
	"strconv"
	"time"
)

func WebsocketInit() {
	Register("login", LoginHandler)
	Register("heartbeat", HeartbeatHandler)
	//Register("ping", PingController)
}

func LoginHandler(client *Client, seq string, message []byte) (code int, msg string, data interface{}) {
	var userInfo out.UserInfo
	err := json.Unmarshal(message, &userInfo)
	if err != nil {
		fmt.Println("Unmarshall userInfo failed. err:", err.Error())
		return 500, "数据解析错误", nil
	}
	fmt.Println("Login Controller, userInfo:", userInfo)
	client.UserId = userInfo.UserID
	client.AppId = userInfo.AppID
	//判断用户登录态
	if !hasLogined(client.UserId) {
		return 302, "用户未登录", nil
	}

	//存储用户相关信息
	fmt.Println("start to save user login info ....")
	service.NewUserService().Login(userInfo)

	//将消息通知到网关
	fmt.Println("LoginUserInfo", userInfo, ", ready to sent grpc, params", seq, userInfo.AppID, userInfo.UserID, userInfo.UserName)

	tmpStruct := protobuf.SendMsgReq{
		Seq:      seq,
		AppId:    uint32(userInfo.AppID),
		UserId:   userInfo.UserID,
		UserName: userInfo.UserName,
		Cmd:      "enter",
		Msg:      "欢迎加入聊天室",
		Type:     "broadcast",
	}

	val, err := proto.Marshal(&tmpStruct)
	if err != nil {
		//todo
		return
	}

	fmt.Println("ready to push to kafka,", string(val))
	common.G_Mq.Push(common.PushMsg{
		Destination: "demo",
		Value:       sarama.ByteEncoder(val),
	})

	fmt.Println("current serverAddr:", config.G_Config.WsAddr, ", lent(clients):", len(G_clientManager.GetCurrentClients()))
	return 0, "用户登录成功", nil
}

func HeartbeatHandler(client *Client, seq string, message []byte)(code int, msg string, data interface{}) {

	client.LastHeartBeat = uint64(time.Now().Unix())

	key := common.RoomUserListRedisPrefixKey + strconv.Itoa(client.AppId)
	_ = common.G_redisClient.HSet(key, client.UserId, time.Now().Format("2006-01-02 15:04:05"))
	return
}


func hasLogined(userId string) bool {
	return true
}
