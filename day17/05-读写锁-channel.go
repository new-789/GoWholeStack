package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ReadGo2(out <-chan int, idx int) {
	for {
		num := <-out // 往 chanel 中读出数据
		fmt.Printf("====== %d th 读 go 程，读出：%d\n", idx, num)
	}
}

func WriteGo2(in chan<- int, idx int) {
	for {
		num := rand.Intn(1000)
		in <- num // 往 channel 中写入数据
		fmt.Printf("%d th 写 go 程，写入：%d\n", idx, num)
		time.Sleep(time.Millisecond * 300)
	}
}

func main0501() {
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())
	// 创建 channel
	ch := make(chan int)

	for i := 0; i < 5; i++ { // 同时创建五个读 go 程
		go ReadGo2(ch, i+1)
	}

	for i := 0; i < 5; i++ { // 同时创建五个写 go 程
		go WriteGo2(ch, i+1)
	}

	for {

	}
}
