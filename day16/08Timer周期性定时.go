package main

import (
	"fmt"
	"runtime"
	"time"
)

func main0801() {
	quit := make(chan bool) // 创建一个 channel 用来终止主 go 程
	fmt.Println("startTime:", time.Now())
	// 创建一个周期定时器，并指定周期时间
	ticker := time.NewTicker(time.Second)

	go func() {
		i := 0
		for {
			// 从 Ticker 结构体中的只读 channel 中读取时间，读不到则一直阻塞
			endTime := <-ticker.C
			i++
			fmt.Println("endTime:", endTime)
			if i == 3 {
				quit <- true     // 解除主 go 程中读取 quit 管道阻塞状态
				runtime.Goexit() // 退出 go 程
			}
		}
	}()

	<-quit // 如果从 quit 管道中读出数据则主 go 程结束，否则一直处于阻塞状态
}
