package main

import (
	"fmt"
	"sync"
	"time"
)

// RWlock 读写锁
var (
	x    = 0
	//lock sync.Mutex // 互斥锁
	rwlock sync.RWMutex // 读写锁
	wg sync.WaitGroup
)

func read() {
	defer wg.Done()
	rwlock.RLock()  // 加读锁
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	rwlock.RUnlock() // 解读锁
}

func write() {
	defer wg.Done()
	rwlock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	rwlock.Unlock()
}

func main() {
	now := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	//time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	end := time.Now()
	wg.Wait()
	fmt.Println(end.Sub(now))
}
