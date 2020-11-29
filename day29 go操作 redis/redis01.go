package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func 2901main() {
	// 连接  redis 操作
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("link redis failed", err)
		return
	}
	defer conn.Close()
	// send 执行 set 增加数据操作命令
	conn.Send("set", "e1", "hello worlds~!")
	if err = conn.Flush(); err != nil {
		fmt.Println("执行命令错误", err)
		return
	}
	fmt.Println(conn.Receive())

	// send 执行 get 操作获取元素
	if err = conn.Send("get", "e1"); err != nil {
		fmt.Println("发送命令错误", err)
		return
	}
	if err = conn.Flush(); err != nil {
		fmt.Println("执行命令错误", err)
		return
	}
	if reply, err := conn.Receive(); err != nil {
		fmt.Println("获取数据错误", err)
		return
	} else {
		fmt.Println(reply)
	}
}
