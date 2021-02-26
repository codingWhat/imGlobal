package out

//登录
type Login struct {
	UserId int `json:"user_id"`
	Source int `json:"source"`
}

//心跳
type HeartBeat struct {
}

//注销
