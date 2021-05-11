package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 原子操作
var x int64 = 0
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	defer wg.Done()
	//lock.Lock()
	atomic.AddInt64(&x, 1)
	//lock.Unlock()
}

func main() {
	wg.Add(100000)
	for i := 0;i < 100000; i++ {
		go add()
	}
	wg.Wait()
	fmt.Println(x)

	// 比较并交换
	//x = 300
	//ok := atomic.CompareAndSwapInt64(&x, x, 200)
	//fmt.Println(ok, x)
}