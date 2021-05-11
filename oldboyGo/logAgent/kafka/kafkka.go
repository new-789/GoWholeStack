package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

// 专门往 kafka 消息队列写日志模块

type logData struct {
	topic string
	data string
}
var (
	client sarama.SyncProducer // 声明一个全局连接 kafka 生产者
	logDataChan chan *logData // 该通到用来接受 tail 包中发送的日志信息
)

// Init 初始化 client
func Init(address []string, maxSize int) (err error) {
	config := sarama.NewConfig()
	// 发送完数据需要 lead 和 follow 都需要确认配置
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 新选出一个 partition
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 成功提交消息在 success channel 返回
	config.Producer.Return.Successes = true

	// 连接 kafka
	client, err = sarama.NewSyncProducer(address, config)
	if err != nil {
		fmt.Printf("producer closed, err:%v\n", err)
		return
	}
	// 初始化 logDataChan
	logDataChan = make(chan *logData, maxSize)
	// 开启一个 goroutine 后台任务一直从 logDataChan 中读去数据并发送到 kafka
	go sendToKafka()
	return nil
}

// sendToKafka 真正发送数据到 kafka 函数
func sendToKafka() {
	for {
		select {
		// 从 logDataChan 中获取数据发送给 kafka
		case logMsg := <-logDataChan:
			// 初始化消息结构
			msg := &sarama.ProducerMessage{}
			msg.Topic = logMsg.topic
			msg.Value = sarama.StringEncoder(logMsg.data)
			// 将消息发送到 kafka
			parId, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Printf("send msg failed, err:%v\n", err)
				return
			}
			fmt.Printf("send data success pid:%v offset:%v value:%s\n", parId, offset, msg.Value)
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}

// SendToChan 暴露给外部的一个函数，用于从 tail 包中将日志发送到内部的 logDataChan 通道中
func SendToChan(topic, data string) {
	logMsg := &logData{
		topic: topic,
		data: data,
	}
	logDataChan <- logMsg
}