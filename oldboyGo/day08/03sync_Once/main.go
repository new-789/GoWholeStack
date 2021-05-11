package main

import (
	"fmt"
	"sync"
)

// sync_Once
var (
	wg sync.WaitGroup
	once sync.Once
)

func f1(ch1 chan <- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1 <- chan int, ch2 chan <- int) {
	defer wg.Done()
	for {
		x, ok := <- ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	// Do 方法接收一个函数做为参数
	once.Do(func() { close(ch2) }) // 保证该操作只执行一次
}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	wg.Add(3)
	go f1(ch1)
	go f2(ch1, ch2)
	go f2(ch1, ch2)
	wg.Wait()
	for ret :=  range ch2 {
		fmt.Println(ret)
	}
}
