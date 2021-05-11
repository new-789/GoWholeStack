package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

// kafka 消费者实例

func main() {
	// 1. 连接 kafka
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Printf("connect to kafka failed, err:%v\n", err)
		return
	}
	// 2. 根据 topic 取得所有的分区
	partitionList, err := consumer.Partitions("web_log")
	if err != nil {
		fmt.Printf("fail to get partition failed, err:%v\n", err)
		return
	}
	fmt.Println("分区列表", partitionList)
	// 3. 遍历所有分区
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer partition %d, err:%v\n", partition, err)
			return
		}
		// defer pc.AsyncClose()

		// 4. 异步从每个分区消费消息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, offset:%d key:%v value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
			}
		}(pc)
	}
	select {}
}
