package ws

import (
	"encoding/json"
	"fmt"
	"github.com/codingWhat/imGlobal/gateway/config"
	"github.com/codingWhat/imGlobal/gateway/data/out"
	"github.com/codingWhat/imGlobal/gateway/servers/grpcclient"
	"github.com/codingWhat/imGlobal/gateway/service"
)

func WebsocketInit() {
	Register("login", LoginHandler)
	//Register("heartbeat",HeartbeatController)
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
	grpcclient.SendMsgAll(seq, userInfo.AppID, userInfo.UserID, userInfo.UserName, "enter", "欢迎加入聊天室")

	fmt.Println("current serverAddr:", config.G_Config.WsAddr, ", lent(clients):", len(G_clientManager.GetCurrentClients()))
	return 0, "用户登录成功", nil
}

func hasLogined(userId string) bool {
	return true
}
