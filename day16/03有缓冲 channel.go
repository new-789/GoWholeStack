package main

import (
	"fmt"
	"time"
)

func main0301() {
	ch := make(chan int, 3) // 创建一个有缓冲 channel 容量为 3
	fmt.Println("len=", len(ch), "cap=", cap(ch))

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Println("子 go 程，i= ", i, "len=", len(ch), "cap=", cap(ch))
		}
	}()

	time.Sleep(time.Second * 2)
	for i := 0; i < 5; i++ {
		num := <-ch
		fmt.Println("主 go 程读到：", num)
	}
}
