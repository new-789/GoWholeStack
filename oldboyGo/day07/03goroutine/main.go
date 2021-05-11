package main

import (
	"fmt"
	"time"
)

// goroutine

// 程序启动之后，会创建一个主 goroutine 去执行
func main() {
	for i := 0; i < 1000; i++ {
		go func(i int) { // 开启一个单独的 goroutine 去执行 hello 函数
			fmt.Println("hello", i)
		}(i)
	}
	fmt.Println("main")
	// main 函数结束了，由 main 函数启动的 goroutine 也都结束了
	time.Sleep(time.Second)
}
