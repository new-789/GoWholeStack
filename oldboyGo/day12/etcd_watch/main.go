package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

// etcd watch demo
func main() {
	// 连接 etcd
	etcdCtl, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Printf("connect etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect etcd success")
	defer etcdCtl.Close()

	// watch 监听操作
	// 派一个哨兵一致监视着 key 为 test 的变化(新增，修改，删除）
	watchCh := etcdCtl.Watch(context.Background(), "test")
	// 从通道尝试取值（监视信息）
	for wresp := range watchCh {
		for _, evt := range wresp.Events {
			fmt.Printf("type:%v key:%v value:%v\n", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))
		}
	}
}
