package main

import (
	"fmt"
	"time"
)

// 定义全局 channel, 用来完成数据同步
var channel = make(chan int)

// 定义一个模拟打印机函数
func printer(s string) {
	for _, v := range s {
		fmt.Printf("%c", v)
		time.Sleep(time.Millisecond * 3000)
	}
}

// 定义两个人使用打印机
func person1() {
	printer("hello")
	channel <- 1 // 给 channel 中加入数据
}

func person2() {
	<-channel // 从 channel 中读数据
	printer("world")
}

func main0601() {
	go person1()
	go person2()

	for {

	}
}
