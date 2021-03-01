package scheduler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/codingWhat/imGlobal/internal/logic/gatecclient"
	"github.com/codingWhat/imGlobal/protobuf"
)

type Handler func(interface{}) error

var jobHandlers map[string] Handler


func RegLogicHandlers() {
	jobHandlers = make(map[string]Handler)

	jobHandlers["single"] = SendMsg
	jobHandlers["broadcast"] = SendMsgAll
}

func getJobHandler(handler string) Handler {
	return jobHandlers[handler]
}


func SendMsg(params interface{}) error {
	ret, _ := json.Marshal(params)
	smR := &protobuf.SendMsgReq{}
	_ = json.Unmarshal(ret, smR)
	_ = gatecclient.SendMsg(smR.Seq, int(smR.AppId), smR.UserId, smR.Cmd, smR.Msg)

	return nil
}

func SendMsgAll(params interface{}) error  {
	ret, _ := json.Marshal(params)
	fmt.Println(ret)
	smR := protobuf.SendMsgReq{}
	_ = json.Unmarshal(ret, &smR)

	fmt.Println("logic Server ready to push to gateway, params:", smR.Seq, int(smR.AppId), smR.UserId, smR.UserName, smR.Cmd, smR.Msg)
	errs := gatecclient.SendMsgAll(smR.Seq, int(smR.AppId), smR.UserId, smR.UserName, smR.Cmd, smR.Msg)

	if len(errs) == 0 {
		return nil
	}

	fmt.Println("broadcast msg failed. err:", errs)
	return errors.New("broadcast msg failed." )
}