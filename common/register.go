package common

//import (
//	"context"
//	"fmt"
//	"github.com/codingWhat/imGlobal/config"
//	"github.com/coreos/etcd/clientv3"
//	"time"
//)
//
//func InitServiceReg() {
//	var (
//		conf clientv3.Config
//		client *clientv3.Client
//		err error
//		putResp *clientv3.PutResponse
//	)
//	conf = clientv3.Config{
//		Endpoints:   []string{config.G_Config.EtcdAddr},
//		DialTimeout: 5 * time.Second,
//	}
//	client, err = clientv3.New(conf)
//	if err != nil {
//		panic(err)
//	}
//
//	kv := clientv3.NewKV(client)
//	putResp, err = kv.Put(context.TODO(), "/servers/" + config.G_Config.GrpcAddr , time.Now().Format("2006-01-02 15:04:05"))
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println("服务注册成功", putResp)
//
//}
//
//func GetAllServerList() {
//	var (
//		conf clientv3.Config
//		client *clientv3.Client
//		err error
//		getResp *clientv3.GetResponse
//	)
//	conf = clientv3.Config{
//		Endpoints:   []string{config.G_Config.EtcdAddr},
//		DialTimeout: 5 * time.Second,
//	}
//	client, err = clientv3.New(conf)
//	if err != nil {
//		panic(err)
//	}
//
//	kv := clientv3.NewKV(client)
//	getResp, err = kv.GetUserInfo(context.TODO(), "/servers/*")
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println("获取所有服务地址", getResp)
//
//}
