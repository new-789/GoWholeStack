package main

import "fmt"

/*
func main0501() {
	ch := make(chan int)  // 双向 chanel

	// 声明一个单向写 channel 并将双向 channel 赋值给该单向 channel
	var sendCh chan <- int = ch
	sendCh <- 888

	// 声明一个单向读 channel 并将双向 channel 赋值给该单向 channel
	var reCh <- chan int = ch
	num := <- reCh
	fmt.Println("num:", num)

	// 反向赋值
	var ch2 chan int = sendCh
}
*/

func send(in chan<- int) {
	in <- 888
	close(in)
}

func recv(out <-chan int) {
	num := <-out
	fmt.Println("读到：", num)
}

func main0501() {
	ch := make(chan int)

	go func() {
		send(ch) // 将双向 channel 传递给 send 函数将其转换为单向写 channel
	}()

	recv(ch) // 将双向 channel 传递给 send 函数将其转换为单向读 channel
}
