package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// WaitGroup
func f() {
	// 初始化随机数种子，保证每次执行的时候都会不一样
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		// 获取随机数
		r1 := rand.Int()
		r2 := rand.Intn(10) // 0 <= x <10
		fmt.Println(r1, r2)
	}
}

var wg sync.WaitGroup

func f1(i int) {
	defer wg.Done() // 计数器减1
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}

func main() {
	// f()
	for i := 0; i < 10; i++ {
		wg.Add(1) // 计数器加1
		go f1(i)
	}
	// 如何知道启动的10个 goroutine 都结束了?
	wg.Wait() // 等待 wg 的计数器减为 0 ，为 0 时表示所有 goroutine 执行完毕
}
