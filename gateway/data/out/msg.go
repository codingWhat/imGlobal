package out

type RequestMsg struct {
	Seq  string      `json:"seq"`
	Cmd  string      `json:"cmd"`
	Data interface{} `json:"data,omitempty"`
}

type UserInfo struct {
	AppID    int    `json:"appId"`
	UserID   string `json:"userId"`
	UserName string `json:"userName"`
}

type BroadMsg struct {
	From string `json:"from"`
}
