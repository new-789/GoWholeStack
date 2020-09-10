package main

import "fmt"

func main0101() {
	ch := make(chan string) // 创建一个屋缓冲 channel
	// len(ch) 求取 channel 中剩余未读取数据个数，cap(ch) 求出 channel 的容量
	fmt.Println("len(ch) = ", len(ch), "cap(ch) = ", cap(ch))
	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println("i =", i)
		}
		// 通知主 go 程打印完毕
		ch <- "子 go 程打印完毕"
	}()

	str := <-ch // 从 channel 中读出数据并存储到一个变量中
	fmt.Println("str = ", str)
}
