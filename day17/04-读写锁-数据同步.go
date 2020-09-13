package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var value int // 定义全局变量，模拟共享数据
// 创建读写锁
var rwmutex1 sync.RWMutex

func ReadGo1(idx int) {
	for {
		// 以读模式加锁，多个 go 程共享共享资源，即多个 go 程可同时使用这把锁
		rwmutex1.RLock()
		num := value
		fmt.Printf("====== %d th 读 go 程，读出：%d\n", idx, num)
		rwmutex1.RUnlock() // 以读模式解锁
	}
}

func WriteGo1(idx int) {
	for {
		// 以写模式加锁，单个 go 程独占共享资源，即多个 go 程只有一个 go 程使用这把锁
		rwmutex1.Lock()
		num := rand.Intn(1000)
		value = num
		fmt.Printf("%d th 写 go 程，写入：%d\n", idx, num)
		time.Sleep(time.Millisecond * 300)
		rwmutex1.Unlock() // 以写模式解锁
	}
}

func main0401() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ { // 同时创建五个读 go 程
		go ReadGo1(i + 1)
	}

	for i := 0; i < 5; i++ { // 同时创建五个写 go 程
		go WriteGo1(i + 1)
	}

	for {

	}
}
