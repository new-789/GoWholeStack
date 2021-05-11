package kafka

import (
	"fmt"
	"github.com/GoWholeStack/oldboyGo/logTransfer/es"
	"github.com/Shopify/sarama"
)

// 初始化 kafka 消费者从 kafka 去取数据发给 ES

// Init 初始化kafka连接
func Init(addr []string, topic string) (err error) {
	// 连接 kafka
	consumer, err := sarama.NewConsumer(addr, nil)
	if err != nil {
		fmt.Printf("connect to kafka failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to kafka success.................")
	// 根据 topic 获取所有分区
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Printf("fail to get list partition err:%v\n", err)
		return err
	}
	fmt.Println("分区列表:", partitionList)

	// 获取本机 ip 当做 es  的类型存入 es
	/*
	ip, err := utils.GetOutboundIP()
	if err != nil {
		fmt.Printf("get local ip failed, err:%v\n", err)
		return err
	}
	 */

	// 遍历所有分区，拿到每一个分区
	for partition := range partitionList {
		// 针对每一个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d, err:%v\n",partition, err)
			return err
		}
		//defer pc.AsyncClose()  # 此处加上会读不出信息

		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, offset:%d key:%v value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
				// 直接发送给 es
				ld := es.LogData{Topic: topic, Data: string(msg.Value)}
				es.SendToESChan(&ld)
			}
		}(pc)
	}
	return
}