package main

import (
	"fmt"
	"sync"
	"time"
)

// 使用 channel 完成多个子 go 程同步
/*
	var ch chan int = make(chan int)

	func printer(str string) {
		for _, v := range str {
			fmt.Printf("%c", v)
			time.Sleep(time.Millisecond * 300)
		}
	}

	func person1() {
		printer("hello")
		ch <- 88  // person1 调用完打印机之后往 channel 中写入数据
	}

	func person2() {
		<-ch  // 在 person2 中从 channel 中读出数据
		printer("world")
	}

	// 使用 channel 完成同步
	func main() {
		go person1()
		go person2()

		for {
			;
		}
	}
*/

// 使用 “锁” 完成数据同步-----互斥锁
var mutex sync.Mutex // 创建一个互斥锁(互斥量)，新建的互斥锁状态为 0，表示未加锁。并且在多个 go 程之间锁只有一把
func printer(s string) {
	mutex.Lock() // 访问共享数据之前，对共享数据进行加锁操作
	for _, ch := range s {
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 300)
	}
	mutex.Unlock() // 访问共享数据结束，对共享数据解锁
}

func person1() {
	printer("hello")
}

func person2() {
	printer("world")
}

func main0201() {
	go person1()
	go person2()

	for {

	}
}
