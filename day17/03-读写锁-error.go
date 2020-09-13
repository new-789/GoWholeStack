package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
) // 1. 导入 sync 标准包

// 创建一个读写锁对象， 无论的读模式还是写模式加解锁都只有这一把锁，但是它有两个属性，分别是读属性和写属性
var rwmutex sync.RWMutex

func ReadGo(out <-chan int, idx int) {
	for {
		rwmutex.RLock() // 通过创建的读写锁对象，以读模式加读锁
		num := <-out
		fmt.Printf("==== %d th 读 go 程，读出：%d\n", idx, num)
		rwmutex.RUnlock() // 通过创建的读写锁对象，以读模式解读锁
	}
}

func WriteGo(in chan<- int, idx int) {
	for {
		// 生成随机数
		num := rand.Intn(100)
		rwmutex.Lock() // 通过创建的读写锁对象，以写模式加写锁
		in <- num
		fmt.Printf("%d th 写 go 程，写入： %d\n", idx, num)
		rwmutex.Unlock()                   // 通过创建的读写锁对象，以写模式加解锁
		time.Sleep(time.Millisecond * 300) // 此行只为放大实验现象
	}
}

func main0301() {
	// 生成随机数种子
	rand.Seed(time.Now().UnixNano())

	//quit := make(chan bool)  // 用于关闭主 go 程 channel
	ch := make(chan int) // 用于数据传递 channel

	for i := 0; i < 5; i++ {
		go ReadGo(ch, i+1)
	}

	for i := 0; i < 5; i++ {
		go WriteGo(ch, i+1)
	}

	//<-quit
	for {

	}
}
