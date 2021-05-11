package main

import (
	"fmt"
	"github.com/GoWholeStack/oldboyGo/logTransfer/conf"
	"github.com/GoWholeStack/oldboyGo/logTransfer/es"
	"github.com/GoWholeStack/oldboyGo/logTransfer/kafka"
	"gopkg.in/ini.v1"
)

// log transfer 将日志从 kafka 取出发往 ES


func main() {
	// 0. 加载配置文件
	var cfg = new(conf.LogTransferCfg)
	err := ini.MapTo(cfg,"./conf/cfg.ini")
	if err != nil {
		fmt.Printf("init config, err:%v\n", err)
		return
	}
	fmt.Printf("cfg: %v\n", cfg)
	// 1. 初始化

	// 1 初始化 ES
	// 1.1:初始化 es 连接的 client
	err = es.Init(cfg.ESCfg.Address, cfg.ESCfg.ChanSize, cfg.ESCfg.MaxGoroutine)
	if err != nil {
		fmt.Printf("init es client failed, err:%v\n", err)
		return
	}

	// 2. 从 kafka 取日志数据
	// 2.1：连接 kafka 创建分区的消费者
	// 2.2：每个分区的消费者分别取出数据，通过 SendToES 函数将数据发往 ES
	err = kafka.Init([]string{cfg.KafkaCfg.Address}, cfg.KafkaCfg.Topic)
	if err != nil {
		fmt.Printf("init kafka Cansumer failed, err:%v\n", err)
		return
	}

	select{}
}
