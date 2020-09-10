package main

import (
	"fmt"
	"time"
)

func main1101() {
	// 创建 ch 管道，用来读写数据
	ch := make(chan int)
	// 创建 quit 管道，用来退出程序
	quit := make(chan bool)

	// 创建 go 程
	go func() {
		for {
			// 通过 for 循环监听 channel 是否有数据流动
			select {
			case num := <-ch:
				fmt.Println("num =", num)
			// 用来监听定时器，以达到 select 超时处理的目的
			case <-time.After(time.Second * 3):
				fmt.Println("timeout")
				// 退出前在子 go 程中向 quit 写入数据，用来通知主 go 程退出
				quit <- true
				/*
					goto 用来跳转到指定位置向下执行代码，通常需要跳转的位置用 label： 来写，
					理论上来说该语法支持跳转到任意位置，但实际上定义的 label: 作用域为局部作用域，所以只能在函数内部进行跳转，
					下面的 goto label 用来跳出 for 循环
				*/
				goto label
				// break 该 break 仅能跳出 select
				//runtime.Goexit()
			}
		}
	label:
		fmt.Println("break goto label---------------")
	}()

	// 在主 go 程中向 ch 管道中写如数据
	for i := 0; i < 2; i++ {
		ch <- i
		time.Sleep(time.Second * 2)
	}

	<-quit // 主 go 程阻塞等待防止主 go 程退出，读取 quit 中的内容后退出程序
	fmt.Println("====== 主 go 程执行完毕 ======")
}
