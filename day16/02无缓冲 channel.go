package main

import (
	"fmt"
	"time"
)

func main0201() {
	ch := make(chan int) // 创建一个无缓冲 channel

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("子 go 程写， i=", i)
			ch <- i
		}
	}()

	time.Sleep(time.Second * 2)
	for i := 0; i < 5; i++ {
		num := <-ch
		fmt.Println("主 go 程读， num=", num)
	}
}
