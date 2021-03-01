package grpcclient

import (
	"context"
	"fmt"
	"github.com/codingWhat/imGlobal/common"
	defs2 "github.com/codingWhat/imGlobal/internal/gateway/defs"
	"github.com/codingWhat/imGlobal/protobuf"
	"google.golang.org/grpc"
	"time"
)

//todo 单发场景对所有网关广播，此处逻辑可以优化处理，将用户和rpc服务地址信息存到redis-cluster，对指定网关进行广播，减少对网关层的tcp压力，但是又增加了一个状态维护(机器下线或者用户离线时处理)
func SendMsg(seq string, appId int, userId, cmd, message string) {

	//获取Gateway机器
	curTime := uint64(time.Now().UnixNano())
	serverList := common.Discovery(curTime)

	for _, addr := range serverList {

		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			fmt.Println("grpc dial 'addr'  failed, err:", err.Error())
			continue
		}

		client := protobuf.NewAccServerClient(conn)
		timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), 100*time.Millisecond)
		//todo
		defer cancelFunc()

		req := protobuf.SendMsgReq{
			Seq:     seq,
			AppId:   uint32(appId),
			UserId:  userId,
			Cmd:     cmd,
			Msg:     message,
			IsLocal: false,
		}

		rsp, err := client.SendMsg(timeoutCtx, &req)
		if err != nil {
			fmt.Println("grpc client SendMsg failed, err:", err.Error())
			continue
		}

		if rsp.RetCode != defs2.RetCodeSuccess {
			fmt.Println("grpc client SendMsg happened bus error, err:", rsp.ErrMsg)
			continue
		}
	}

}

func SendMsgAll(seq string, appId int, userId, userName, cmd, message string) {
	//获取Gateway机器
	curTime := uint64(time.Now().UnixNano())
	serverList := common.Discovery(curTime)

	for _, addr := range serverList {

		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			fmt.Println("grpc dial 'addr'  failed, err:", err.Error())
			continue
		}

		rpcClient := protobuf.NewAccServerClient(conn)
		timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), 100*time.Millisecond)
		//todo
		defer cancelFunc()

		req := protobuf.SendMsgAllReq{
			Seq:      seq,
			AppId:    uint32(appId),
			UserId:   userId,
			UserName: userName,
			Cmd:      cmd,
			Msg:      message,
		}
		rsp, err := rpcClient.SendMsgAll(timeoutCtx, &req)
		if err != nil {
			fmt.Println("grpc client SendMsgAll failed, err:", err.Error())
			continue
		}
		fmt.Println("[*******]msg:", req.Msg, ", sent to server successfully:", addr, "[*******]")

		if rsp.RetCode != defs2.RetCodeSuccess {
			fmt.Println("grpc client SendMsgAll happened bus error, err:", rsp.ErrMsg)
			continue
		}
	}
}
