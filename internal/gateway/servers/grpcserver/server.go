package grpcserver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/codingWhat/imGlobal/common"
	config2 "github.com/codingWhat/imGlobal/internal/gateway/config"
	out2 "github.com/codingWhat/imGlobal/internal/gateway/data/out"
	ws2 "github.com/codingWhat/imGlobal/internal/gateway/servers/ws"
	"github.com/codingWhat/imGlobal/protobuf"
	"google.golang.org/grpc"
	"net"
)

func StartGrpcServer() {
	lis, err := net.Listen("tcp", config2.G_Config.GrpcAddr)
	if err != nil {
		panic(err)
	}
	newServer := grpc.NewServer()
	protobuf.RegisterAccServerServer(newServer, &server{})

	if err = newServer.Serve(lis); err != nil {
		fmt.Println("grpc server call Serve() failed, err:", err.Error())
		panic(err)
	}
}

type server struct {
}

func (s *server) SendMsg(ctx context.Context, reqMsg *protobuf.SendMsgReq) (res *protobuf.SendMsgRsp, err error) {

	for client := range ws2.G_clientManager.Clients {
		if client.UserId == reqMsg.UserId && uint32(client.AppId) == reqMsg.AppId {

			fmt.Println("grpcServer SendMsg recv info:", reqMsg.AppId, reqMsg.UserId, reqMsg.Cmd, reqMsg.Seq, reqMsg.Msg)
			resp := out2.NewResponseDataPack(reqMsg.Seq, reqMsg.Cmd, &common.Response{
				Code: 0,
				Msg:  "ok",
				Data: map[string]string{
					"from": reqMsg.UserName,
					"msg":  reqMsg.Msg,
				},
			})

			ret, err := json.Marshal(resp)
			if err != nil {
				return &protobuf.SendMsgRsp{
					RetCode:   500,
					ErrMsg:    err.Error(),
					SendMsgId: reqMsg.Seq,
				}, nil
			}
			client.SendChan <- ret
		}

	}

	return
}

func (s *server) SendMsgAll(ctx context.Context, reqMsg *protobuf.SendMsgAllReq) (res *protobuf.SendMsgAllRsp, err error) {

	//todo 需要对指定房间下的链接进行推送

	for client := range ws2.G_clientManager.Clients {

		if uint32(client.AppId) != reqMsg.AppId || client.UserId == reqMsg.UserId {
			continue
		}

		fmt.Println("grpcServer SendMsgAll recv info:", uint32(client.AppId), reqMsg.AppId, reqMsg.UserId, reqMsg.Cmd, reqMsg.Seq, reqMsg.Msg)

		resp := out2.NewResponseDataPack(reqMsg.Seq, reqMsg.Cmd, &common.Response{
			Code: 0,
			Msg:  "ok",
			Data: map[string]string{
				"from": reqMsg.UserId + "-" + reqMsg.UserName,
				"msg":  reqMsg.Msg,
			},
		})

		ret, err := json.Marshal(resp)
		if err != nil {
			return &protobuf.SendMsgAllRsp{
				RetCode:   500,
				ErrMsg:    err.Error(),
				SendMsgId: reqMsg.Seq,
			}, nil
		}
		client.SendChan <- ret
	}

	return &protobuf.SendMsgAllRsp{
		RetCode:   0,
		ErrMsg:    "Ok",
		SendMsgId: reqMsg.Seq,
	}, nil

}
