package main

import (
	"fmt"
	"sync"
)

// channel 练习
// 1. 启动一个 groutine，生成100个数发送到 ch1
// 2. 启动一个 goroutine，从 ch1 中取值计算其平方并放在 ch2 中
// 3. 在 main 函数中冲 ch2 中取值并打印
var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for {
		v, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- v * v
	}
	once.Do(func() { close(ch2) }) // 确保某个操作只执行一次
}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	wg.Add(3)
	go f1(ch1)
	go f2(ch1, ch2)
	go f2(ch1, ch2)
	wg.Wait()
	for v := range ch2 {
		fmt.Println(v)
	}
}
