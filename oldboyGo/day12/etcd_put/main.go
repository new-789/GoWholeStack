package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

// etcd Put

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
	value := `[{"path":"/home/zf/logagen_test/nginx.log","topic":"web_log"},{"path":"/home/zf/logagen_test/redis.log","topic":"redis_log"}]`
	//value := `[{"path":"/home/zf/logagen_test/nginx.log","topic":"web_log"},{"path":"/home/zf/logagen_test/redis.log","topic":"redis_log"},{"path":"/home/zf/logagen_test/mysql.log","topic":"mysql_log"}]`
	// 发送数据给 etcd
	_, err = ctl.Put(ctx, "/logagent/192.168.8.100/collect_config", value) // %s 为本机对外 ip
	cancel()
	if err != nil {
		fmt.Printf("put data to etcd failed, err:%v\n", err)
		return
	}
}
