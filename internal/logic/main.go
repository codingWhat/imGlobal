package main

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/codingWhat/imGlobal/common"
	"github.com/codingWhat/imGlobal/internal/logic/config"
	"github.com/codingWhat/imGlobal/internal/logic/scheduler"
	"github.com/codingWhat/imGlobal/protobuf"
)

func main() {
	//初始化配置
	config.InitConfig()

	//初始化服务发现
	common.InitDiscovery()
	if config.G_Config.Topic == "" || config.G_Config.Group == "" {
		panic("consumer topic/group empty..")
	}

	//初始化任务调度器
	scheduler.InitScheduler()

	//启动MQ
	common.InitMq()
	common.G_Mq.StartPull(config.G_Config.Topic, config.G_Config.Group, func(part int32, msg *sarama.ConsumerMessage, partitionOffsetManager sarama.PartitionOffsetManager) {
		worker := scheduler.G_scheduler.GetWorker()
		req := &protobuf.SendMsgReq{}
		_ = json.Unmarshal(msg.Value, req)

		worker.JobChan <- &scheduler.Job{
			Handler: req.Type,
			Params:  req,
		}

		select {
		case rs := <-worker.RetChan:
			if rs {
				fmt.Println("partition:", part, ", topic:", msg.Topic, ",val:", string(msg.Value))
				nextOffset, offsetString := partitionOffsetManager.NextOffset()
				fmt.Println(".... offsetString", offsetString)
				partitionOffsetManager.MarkOffset(nextOffset+1, "modified metadata")
			}
		}
	})

	select {}

	//信号-handler
	//signal.Notify()
}
