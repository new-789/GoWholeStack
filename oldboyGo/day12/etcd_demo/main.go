package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

// etcd demo

func main() {
	// 连接 etcd
	ctl, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Printf("create etcd client failed, err:%v\n", err)
		return
	}
	fmt.Println("connect etcd success!")
	defer ctl.Close()

	// 创建一个 ctx 和 cancel 指定超时时间
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	// 发送数据给 etcd
	_, err = ctl.Put(ctx, "test", "这是一个测试")
	cancel()
	if err != nil {
		fmt.Printf("put data to etcd failed, err:%v\n", err)
		return
	}

	// 从 etcd 中获取数据
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	// clientv3.WithPrefix() 表示使用前缀的方式去除所有的内容，选传参数可不传递
	resp, err := ctl.Get(ctx, "test", clientv3.WithPrefix())
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, v := range resp.Kvs {
		fmt.Printf("Key:%v Value:%v\n", string(v.Key), string(v.Value))
	}
}
