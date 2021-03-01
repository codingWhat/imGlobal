package user

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	models2 "github.com/codingWhat/imGlobal/api/models"
	"github.com/codingWhat/imGlobal/common"
	"github.com/codingWhat/imGlobal/protobuf"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetList(ctx *gin.Context) {
	var (
		appId string
		err   error
		ret   map[string][]string
	)

	appId = ctx.DefaultQuery("appId", "101")
	ret, err = models2.NewUserModel().GetRoomUsers(appId)
	if err != nil {
		fmt.Println(err)
		common.NewResponse(common.CodeSysError, err.Error(), "").Send(ctx)
		return
	}

	common.NewResponse(common.CodeSuccess, "OK", ret).Send(ctx)
}

func SendMsgAll(ctx *gin.Context) {
	var (
		appId    string
		userId   string
		msgId    string
		message  string
		userName string
		err      error
	)

	appId = ctx.DefaultPostForm("appId", "101")
	userId = ctx.DefaultPostForm("userId", "")
	msgId = ctx.DefaultPostForm("msgId", "")
	message = ctx.DefaultPostForm("message", "")

	iAppId, err := strconv.Atoi(appId)
	if err != nil {
		common.NewResponse(common.CodeSysError, err.Error(), "").Send(ctx)
		return
	}
	userName, err = models2.GetUserInfo(userId)
	if err != nil {
		common.NewResponse(common.CodeSysError, err.Error(), "").Send(ctx)
		return
	}


	tmpStruct := protobuf.SendMsgReq{
		Seq:     msgId,
		AppId:   uint32(iAppId),
		UserId:  userId,
		UserName: userName,
		Cmd:     "msg",
		Msg:     message,
		IsLocal: false,
		Type: "broadcast",
	}
	val, _ := json.Marshal(tmpStruct)

	common.G_Mq.Push(common.PushMsg{
		Destination: "demo",
		Value: sarama.ByteEncoder(val),
	})


	common.NewResponse(common.CodeSuccess, "OK", "").Send(ctx)
}

func SendMsg(ctx *gin.Context) {

}
