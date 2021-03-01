package common

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

var G_Mq *Kafka

func InitMq() {
	G_Mq = NewKafka()
}

type Kafka struct {
	asyncProducer sarama.AsyncProducer
	syncProducer  sarama.SyncProducer
	consumer      sarama.Consumer

}

func NewKafka() *Kafka {
	kafka := &Kafka{}
	kafka.initConsumer()
	kafka.initProducer()
	return kafka
}

func (k *Kafka) initProducer() {
	cfg := sarama.NewConfig()

	cfg.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	cfg.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	cfg.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms

	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, cfg)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	go func() {
		for err := range producer.Errors() {
			log.Println("Failed to write access log entry:", err)
		}
	}()

	k.asyncProducer = producer
}

func (k *Kafka) initConsumer() {

	cfg := sarama.NewConfig()
	cfg.Consumer.Offsets.CommitInterval = 1 * time.Second
	cfg.Version = sarama.V0_10_0_1

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	k.consumer = consumer
}

func (k *Kafka) Push(msg PushMsg) {
	//todo 池化这个消息对象
	fmt.Println(sarama.ProducerMessage{
		Topic: msg.Destination,
		Value: sarama.ByteEncoder(msg.Value),
	})
	k.asyncProducer.Input() <- &sarama.ProducerMessage{
		Topic: msg.Destination,
		Value: sarama.ByteEncoder(msg.Value),
	}

	fmt.Println("sent to kafka successfully...")
}

func (k *Kafka) StartPull(topic  string, group string, handler func(int32, *sarama.ConsumerMessage, sarama.PartitionOffsetManager)) {
	partitions,_:= k.consumer.Partitions(topic)
	fmt.Println(partitions)

	defer func() {
		if err := k.consumer.Close(); err != nil {
			fmt.Println("consumer close err:", err.Error())
		}
	}()

	for _, part := range partitions {
		go func(part int32) {
			cfg := sarama.NewConfig()
			cfg.Consumer.Offsets.CommitInterval = 1 * time.Second
			cfg.Version = sarama.V0_10_0_1
			client, err := sarama.NewClient([]string{"localhost:9092"}, cfg)
			if err != nil {
				fmt.Println(err)
				return
			}

			defer func() {
				_ = client.Close()
			}()
			offsetManager, err := sarama.NewOffsetManagerFromClient(group, client)
			partitionOffsetManager, err := offsetManager.ManagePartition(topic, part)
			defer func() {
				_ = partitionOffsetManager.Close()
			}()

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			nextOffset,_:= partitionOffsetManager.NextOffset()
			fmt.Println(nextOffset)

			partitionConsumer, err := k.consumer.ConsumePartition(topic, part, nextOffset+1)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			defer func() {
				_ = partitionConsumer.Close()
			}()

			for {
				select {
				  case msg := <-partitionConsumer.Messages():
					  handler(part, msg, partitionOffsetManager)
				}
			}

		}(part)

	}

	select {

	}

	//outChan := make(chan os.Signal)
	//signal.Notify(outChan, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGQUIT)
	//select {
	//case  sig := <-outChan:
	//}
}
