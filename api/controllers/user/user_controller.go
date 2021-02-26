package user

import (
	"fmt"
	"github.com/codingWhat/imGlobal/api/models"
	"github.com/codingWhat/imGlobal/common"
	"github.com/codingWhat/imGlobal/gateway/servers/grpcclient"
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
	ret, err = models.NewUserModel().GetRoomUsers(appId)
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
	userName, err = models.GetUserInfo(userId)
	if err != nil {
		common.NewResponse(common.CodeSysError, err.Error(), "").Send(ctx)
		return
	}

	grpcclient.SendMsgAll(msgId, iAppId, userId, userName, "msg", message)

	common.NewResponse(common.CodeSuccess, "OK", "").Send(ctx)
}

func SendMsg(ctx *gin.Context) {

}
