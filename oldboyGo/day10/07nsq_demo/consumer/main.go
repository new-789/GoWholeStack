package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// NSQ_CONSUMER Demo
type MyHandler struct {
	Title string
}

// HandleMessages 实现一个消费者方法
func (m *MyHandler)HandleMessage(msg *nsq.Message) (err error) {
	fmt.Printf("%s recv from %v, msg:%v\n", m.Title, msg.NSQDAddress, string(msg.Body))
	return
}

// 初始化消费者
func initConsumer(topic, channel, address string) (err error) {
	config := nsq.NewConfig() // 初始化配置信息
	// 设置 15 秒去 lookupd 中查询一次数据
	config.LookupdPollInterval = 15 * time.Second
	// 初始化一个消费者
	c, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		fmt.Printf("create cunsumer failed, err:%v\n", err)
		return
	}
	consumer := &MyHandler{
		Title:"天河1号",
	}
	c.AddHandler(consumer)

	// 通过 lookupd 查询
	// if err := ConnectToNSQD(address); err != nil { //直接连 NSQD
	if err := c.ConnectToNSQLookupd(address); err != nil {
		return err
	}
	return nil
}

func main() {
	err := initConsumer("topic_demo", "first", "127.0.0.1:4161")
	if err != nil {
		fmt.Printf("init consumer failed, err:%v\n", err)
		return
	}

	c := make(chan os.Signal) // 定义一个信号通道
	signal.Notify(c, syscall.SIGINT) // 转发键盘中断信号到 c
	<- c // 阻塞
}
