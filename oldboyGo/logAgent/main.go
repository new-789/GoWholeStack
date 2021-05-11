package main

import (
	"fmt"
	"github.com/GoWholeStack/oldboyGo/logAgent/conf"
	"github.com/GoWholeStack/oldboyGo/logAgent/etcd"
	"github.com/GoWholeStack/oldboyGo/logAgent/kafka"
	"github.com/GoWholeStack/oldboyGo/logAgent/taillog"
	"github.com/GoWholeStack/oldboyGo/logAgent/utils"
	"gopkg.in/ini.v1"
	"sync"
	"time"
)

var (
	// 因为 conf.AppConfig 是指针类型，所以在定义全局变量时需要使用 new 方法初始化
	cfg = new(conf.AppConfig)
)

/*
func run() {
	// 0. 读取配置文件
	topic := cfg.KafkaConfig.Topic
	// 1. 读取日志
	for {
		// 2. 发送到 kafka
		select {
		// 从 tail channel 中读取数据
		case line := <-taillog.ReadChan():
			// 发送数据
			kafka.SendToKafka(topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}
*/

// logAgent 入口程序
func main() {
	// 0. 加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}
	address := cfg.KafkaConfig.Address
	// 1. 初始化 kafka 文件
	err = kafka.Init([]string{address}, cfg.KafkaConfig.ChanMaxSize)
	if err != nil {
		fmt.Printf("init kafka failed, err:%v\n", err)
		return
	}
	fmt.Println("init kafka success!")
	// 2. 初始化 etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("init etcd success!")
	// 为了实现每个 logAgent 都拉取自己独有的配置，所以要以自己的 ip 地址作为区分
	ipStr, err := utils.GetOutboundIP()
	if err != nil {
		panic(err)
	}
	etcdConfKey := fmt.Sprintf(cfg.EtcdConf.Key, ipStr)
	// 2.1 从 etcd 中获取日志收集项的配置信息
	logEntryConf, err := etcd.GetConf(etcdConfKey)
	if err != nil {
		fmt.Printf("getconf failed, err:%v\n", err)
		return
	}
	fmt.Printf("======get conf from etcd success, :%v\n", logEntryConf)

	for i, v := range logEntryConf {
		fmt.Printf("index:%v value:%v\n", i, v)
	}

	// 3. 收集日志发往 kafka
	// 因为 NewConfChan 访问了 taillog 模块中的 tailObjMgr 的 NewConfChan，这个 channel 是在 taillog.Init(logEntryConf) 执行的
	taillog.Init(logEntryConf)

	// 2.2 派一个哨兵在后台运行实时监视日志收集项的变化(有变化及时通知我的 logagent 实现热加载）
	newConfChan := taillog.NewConfChan() // 从 taillog 模块中获取对外暴露的通道
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(etcdConfKey, newConfChan) // 哨兵发现发现最新的配置信息会通知上面 newConfChan 通道
	wg.Wait()

	// 3. 具体的业务逻辑，往 kafka 发送数据
	//run()
}
