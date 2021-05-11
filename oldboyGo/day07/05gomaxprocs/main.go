package main

import (
	"fmt"
	"runtime"
	"sync"
)

//GOMAXPROCS
var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("A", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("B", i)
	}
}

func main() {
	runtime.GOMAXPROCS(2)         // 设置 CPU 核心数
	fmt.Println(runtime.NumCPU()) // 返回 cpu 逻辑核心数
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
