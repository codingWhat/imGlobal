package common

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func (r *Response) Send(ctx *gin.Context) {
	res, err := json.Marshal(r)
	if err != nil {
		fmt.Println("response marshall response failed. err:", err.Error())
	}
	n, err := ctx.Writer.Write(res)
	if err != nil {
		fmt.Println("response write failed. err:", err.Error())
	} else {
		fmt.Println("response write success. n:", n)
	}
}

func NewResponse(code int, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
