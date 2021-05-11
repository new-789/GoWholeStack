package main

import (
	"fmt"
)

// channel
var ch chan int // 需要指定通道中元素的类型
// var wg sync.WaitGroup

func noBufchannel(ch chan int) {
	ch = make(chan int) // 不带缓冲区的通道初始化
	go func() {
		x := <-ch // 从通道中获取值
		fmt.Println("从通道中取到了：", x)
	}()
	ch <- 10 // 往通道中发送值
	fmt.Println("10 发送到到了 channel 通道中了......")
}

func bufChannel(ch chan int) {
	fmt.Println(ch)
	ch = make(chan int, 1) // 带缓冲去的通道初始化
	ch <- 10               // 往通道中发送值
	fmt.Println("10 发送到到了 channel 通道中了......")
	ch <- 20
	fmt.Println("20 发送到到了 channel 通道中了......")
	x := <-ch // 从通道中获取值
	fmt.Println("从通道中取到了：", x)
	close(ch) // 关闭通道
}

func main() {
	bufChannel(ch)
}
