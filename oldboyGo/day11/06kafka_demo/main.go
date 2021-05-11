package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// kafka demo

// sarama 包使用
func main() {
	// 初始化配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // 发完数据需要 leader 和 follow 都确认模式
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个 partition
	config.Producer.Return.Successes = true // 成功交付的消息将在 success_channel 返回
	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log" // 指定 topic 名称
	// 将发送的值序列化
	msg.Value = sarama.StringEncoder("this is a test log")
	// 连接 kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Printf("producer closed, err:%v\n", err)
		return
	}
	defer client.Close()
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed, err:", err)
		return
	}
	fmt.Printf("====== pid:%v offset:%v\n", pid, offset)
}
