package main

import (
	"fmt"
	"runtime"
)

func main0301() {

	go func() {
		for {
			fmt.Println("this is goroutine test")
		}
	}()

	for {
		runtime.Gosched() // 出让当前 cpu 时间轮片
		fmt.Println("------ i'm is main ------")
	}
}
