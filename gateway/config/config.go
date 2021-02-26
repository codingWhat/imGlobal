package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var G_Config *Config

func InitConfig() {
	//加载配置文件
	data, err := ioutil.ReadFile("./config/main.json")
	if err != nil {
		fmt.Println("read config file failed. err:", err.Error())
		panic(err)
	}
	G_Config = NewConfig()
	err = G_Config.FromJson(data)
	if err != nil {
		fmt.Println("config file json unmarshal failed. err:", err.Error())
		panic(err)
	}
}

type Config struct {
	WsAddr          string `json:"wsServerAddr"`
	WsReadBuffSize  int    `json:"wsServerReadBuffSize"`
	WsWriteBuffSize int    `json:"wsServerWriteBuffSize"`
	HttpAddr        string `json:"httpServerAddr"`
	EtcdAddr        string `json:"etcdAddr"`
	GrpcAddr        string `json:"grpcServerAddr"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) FromJson(data []byte) error {
	return json.Unmarshal(data, c)
}
