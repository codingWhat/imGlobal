package out

import "github.com/codingWhat/imGlobal/common"

func NewResponseDataPack(seq string, cmd string, response *common.Response) *ResponseDataPack {
	return &ResponseDataPack{
		Seq:      seq,
		Cmd:      cmd,
		Response: response,
	}
}

type ResponseDataPack struct {
	Seq      string           `json:"seq"`
	Cmd      string           `json:"cmd"`
	Response *common.Response `json:"response,omitempty"`
}
