package main

import (
	"fmt"
	"sync"
	"time"
)

// 为什么需要 context
var wg sync.WaitGroup
var notify bool

func f() {
	defer wg.Done()
	for {
		fmt.Println("hello")
		time.Sleep(time.Millisecond * 500)
		if notify {
			break
		}
	}
}

func main() {
	wg.Add(1)
	go f()
	time.Sleep(5 * time.Second)
	notify = true // 表示通知子 goroutine 该退出了
	wg.Wait()
	// 如何通知子 goroutine 提出
}
