package main

import (
	"fmt"
	"time"

	"github.com/gitstliu/go-redis-cluster"
	// "github.com/gomodule/redigo/redis"
)

func main() {
	// 连接集群操作,返回一个连接对象，和错误信息
	cluster, err := redis.NewCluster(
		// &redis.Options 该参数用来设置集群的配置文件 []string 类型，其中的元素便是每个配置文件中绑定的 ip:port
		&redis.Options{
			StartNodes: []string{
				"192.168.8.100:7001",
				"192.168.8.100:7002",
				"192.168.8.100:7003",
				"192.168.8.100:7004",
				"192.168.8.100:7005",
				"192.168.8.100:7006",
			},
			ConnTimeout:  50 * time.Microsecond, // 链接超时时间
			ReadTimeout:  50 * time.Microsecond, // 读取超时事件
			WriteTimeout: 50 * time.Microsecond, // 写入超时时间
			KeepAlive:    16,
			AliveTime:    60 * time.Second,
		})
	// 执行 redis 命令
	rel, err := cluster.Do("set", "hello", "world")
	if err != nil {
		fmt.Println("操作 redis 错误", err)
		return
	}
	fmt.Println(rel)
}
