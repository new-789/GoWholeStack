package main

import (
	"fmt"
	"sync"
	"time"
)

// 为什么需要 context
var wg sync.WaitGroup
var c = make(chan int, 1)

func f() {
	defer wg.Done()
	LOOP:
		for {
			fmt.Println("hello")
			time.Sleep(time.Millisecond * 500)
			select {
			case <-c:
				break LOOP // 指定跳到指定位置
			default:
			}
		}
}

func main() {
	wg.Add(1)
	go f()
	time.Sleep(5 * time.Second)
	c <- 1
	wg.Wait()
	// 如何通知子 goroutine 提出
}
