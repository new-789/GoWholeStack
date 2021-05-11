package main

import (
	"fmt"
	"sync"
)

// 锁
var (
	x = 0
	wg sync.WaitGroup
	lock sync.Mutex
)

func add() {
	for i := 0; i < 50000; i++ {
		lock.Lock() //  加互斥锁操作
		x = x + 1
		lock.Unlock() // 解互斥锁操作
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
