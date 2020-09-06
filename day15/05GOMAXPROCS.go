package main

import (
	"fmt"
	"runtime"
)

func main0501() {
	// 设置 cpu 最核心数为 1 来执行程序，并返回上一次设置的值，由于之前没有设置则为默认 cpu 核心数
	n := runtime.GOMAXPROCS(1)
	fmt.Println("n =", n)
	// 设置 cpu 最核心数为 2 来执行程序，并返回上一次设置的值 1
	n = runtime.GOMAXPROCS(2)
	fmt.Println("n =", n)
	for {
		go fmt.Print(0) // 子 go 程
		fmt.Print(1)    // 主 go 程
	}
}
