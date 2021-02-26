package home

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Index(c *gin.Context) {
	appIdStr := c.Query("appId")
	appIdUint64, _ := strconv.ParseInt(appIdStr, 10, 32)
	appId := uint32(appIdUint64)

	fmt.Println("http_request 聊天首页", appId)

	data := gin.H{
		"title":        "聊天首页",
		"appId":        appId,
		"httpUrl":      "iml.server.com",
		"webSocketUrl": "iml.server.com",
		//"httpUrl":      "192.168.124.13:8080",
		//"webSocketUrl": "192.168.124.13:8089",
	}
	c.HTML(http.StatusOK, "index.tpl", data)
}
