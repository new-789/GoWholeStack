package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

// LogEntry 需要收集的日志配置信息
type LogEntry struct {
	Path string `json:"path"`  // 日志存放的路径
	Topic string `json:"topic"` // 日志要发往 kafka 中的 Topic
}

var (
	cli *clientv3.Client
)

// Init 初始化 etcd 函数
func Init(addr string, timeout time.Duration) (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints: []string{addr},
		DialTimeout: timeout,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	return
}

// GetConf 从 etcd 中根据 key 获取配置项
func GetConf(key string) (logEntry []*LogEntry, err error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, v := range resp.Kvs {
		// 将 json 格式数据反序列化到 logEntry 中
		err = json.Unmarshal(v.Value, &logEntry)
		if err != nil {
			fmt.Printf("unmarshal etcd value failed, err:%v\n", err)
			return
		}
	}
	return
}

// WatchConf 开启哨兵实时监视日志收集项目配置内容，如有配置更新则通知 tailTask 任务进行相应调整
func WatchConf(key string, newConfChan chan <- []*LogEntry) {
	ch := cli.Watch(context.Background(), key)
	for wresp := range ch {
		for _, evt := range wresp.Events {
			// 通知 taillog.tailObjMgr
			// 1. 先判断操作的类型
			var newConf []*LogEntry
			if evt.Type != clientv3.EventTypeDelete {
				// 如果不是删除操作，则通知 tailTask 新的配置来了
				err := json.Unmarshal(evt.Kv.Value, &newConf)
				if err != nil {
					fmt.Printf("json unmarshal config info failed, err:%v\n", err)
					continue
				}
			}
			fmt.Println("--------------------get newConf:", newConf)
			newConfChan <- newConf
		}
	}
}