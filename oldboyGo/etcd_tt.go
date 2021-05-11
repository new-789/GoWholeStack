package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func main() {
	clt, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	// 设置续期 5 秒
	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}
	// 设置 k-v 到 etcd
	_, err = cli.Put(context.TODO(), "root", "admin", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}
	// 若想一直有效，设置自动续期
	ch, err := cli.KeepAlive(context.TODO(), resp.ID)
	if err != nil {
		log.Fatal(err)
	}
	for {
		c := <- ch
		fmt.Println(c)
	}
}
