package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 在全局创建条件变量
var cond sync.Cond

// 生产者函数
func producer1(in chan<- int, idx int) {
	for {
		cond.L.Lock()            // 使用条件变量对生产者进行加互斥锁操作
		for len(in) == cap(in) { // 使用 for 循环判断公共区是否已满
			cond.Wait() // 挂起当前 go 程，等到条件变量满足，被消费者唤醒
		}
		num := rand.Intn(1000)
		in <- num // 往 channel 中写入数据
		fmt.Printf("%d th 生产者，生产了：%d\n", idx, num)
		cond.L.Unlock()         // 生产结束，使用条件变量进行解锁操作
		cond.Signal()           // 唤醒公共区对端阻塞的消费者
		time.Sleep(time.Second) // 生产一个数据休息一秒
	}
}

// 消费者函数
func consumer1(out <-chan int, idx int) {
	for {
		cond.L.Lock()       // 使用条件变量对消费者进行加互斥锁操作
		for len(out) == 0 { // 是否 for 循环判断公共区是否已空，空则阻塞当前 go 程
			cond.Wait() // 挂起当前 go 程，等待条件变量满足，被生产者唤醒
		}
		num := <-out // 从 channel 中读出数据
		fmt.Printf("%d th 消费者，消费了：%d\n", idx, num)
		cond.L.Unlock() // 消费结束，使用条件变量进行解锁操作
		cond.Signal()   // 唤醒公共区对端的生产者
	}
}

func main0701() {
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	// 创建有缓冲 channel 模拟共享数据
	ch := make(chan int, 3)

	// 为条件变量中的 L 成员指定锁的类型，此处指定的为互斥锁
	cond.L = new(sync.Mutex)

	// 创建消费者 go 程
	for i := 0; i < 5; i++ {
		go consumer1(ch, i+1)
	}

	// 创建生产者 go 程
	for i := 0; i < 3; i++ {
		go producer1(ch, i+1)
	}

	for {

	}
}
