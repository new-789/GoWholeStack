package main

import (
	"fmt"
	"time"
)

func main0701() {
	fmt.Println("当前时间:", time.Now())
	// 创建 timer 定时器对象，并指定等待的时间
	timer := time.NewTimer(time.Second * 2)
	// 从 timer 定时器中读取时间
	nowTime := <-timer.C
	fmt.Println("现下时间:", nowTime)
}

// 三种定时方法
func main0702() {
	// 1. Sleep
	/*
		time.Sleep(time.Second)
	*/

	// 2. timer.C
	/*
		// 创建 timer 定时器对象，并指定定时时长
		timer := time.NewTimer(time.Second * 2)
		// 从 timer 定时器中读取时间，定时满系统自动写入系统当前时间
		nowTime := <- timer.C
		fmt.Println("现下时间:", nowTime)
	*/

	// 3. time.After
	fmt.Println("当前时间:", time.Now())
	nowTime := <-time.After(time.Second * 2)
	fmt.Println("现下时间:", nowTime)
}

// 定时器的停止和重置
func main0703() {
	timer := time.NewTimer(time.Second * 10) // 创建定时器，并设定定时时长

	timer.Reset(time.Second * 1) // 重置定时时长为 1
	go func() {
		<-timer.C // 读取定时器中的时间，如果定时器停止此处则会阻塞
		fmt.Println("子 go 程定时完毕.....")
	}()

	//timer.Stop()   // 设置定时器停止
	for {

	}
}
