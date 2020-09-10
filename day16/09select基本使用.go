package main

import (
	"fmt"
	"runtime"
	"time"
)

func main0901() {
	ch := make(chan int)    // 创建 channel 用来进行数据通信的
	quit := make(chan bool) // 用来判断是否退出的 channel

	go func() { // 子 go 程写数据
		for i := 0; i < 5; i++ {
			ch <- i // 往 ch 管道中写入数据
			time.Sleep(time.Second)
		}
		close(ch)    // 关闭 channel
		quit <- true // 通知主 go 程退出
		runtime.Goexit()
	}()

	for { // 注 Go 程读数据
		select {
		case num := <-ch: // 监听 ch 管道，如果有数据则从中取出数据
			fmt.Println("主 go 程读到：", num)
		case <-quit: // 监听 quit 管道，如果读到数据则退出程序
			fmt.Println("======= 主 Go 程执行结束 ======")
			//break  写在 select 语句中可用来跳出 select 语句
			//runtime.Goexit()  该方法只能用来退出子 go 程
			return
		}
	}
}
