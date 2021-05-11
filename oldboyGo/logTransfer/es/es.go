package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
	"time"
)

// 初始化 es ，准备接受 kafka 发来的数据

// LogData ...
type LogData struct {
	Topic string `json:"topic"`
	Data string `json:"data"`
}

var (
	client *elastic.Client
	ch chan *LogData
)

// Init ...
func Init(addr string, chanSize, MaxGoroutine int) (err error) {
	if !strings.HasSuffix(addr, "http://") {
		addr = "http://" + addr
	}
	client, err = elastic.NewClient(elastic.SetURL(addr))
	if err != nil {
		fmt.Printf("connect to elastic failed. err:%v\n", err)
		return
	}
	fmt.Println("connect to es success")
	// 初始化 es 完成后开启 MaxGoroutine 个进程在后台从 channel 中获取数据发给 es
	ch= make(chan *LogData, chanSize)
	for i := 0; i < MaxGoroutine; i++ {
		go SendToES()
	}
	return
}

// SendToESChan 发送数据到 channel
func SendToESChan(msg *LogData) {
	ch <- msg
}

// SendToES 发送数据到 elasticsearch
func SendToES() {
	for {
		select {
		case msg := <- ch:
			put, err := client.Index().Index(msg.Topic).BodyJson(msg).Do(context.Background())
			if err != nil {
				fmt.Printf("send to es failed, err:%v\n", err)
				continue
			}
			fmt.Println("send to es success ===========",put.Id, put.Index, put.Type)
		default:
			time.Sleep(time.Second)
		}
	}
}