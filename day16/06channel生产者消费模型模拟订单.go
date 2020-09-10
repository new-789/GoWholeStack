package main

import "fmt"

// 创建一个结构体模拟订单的信息
type OrderInfo struct {
	id int
}

// 创建一个函数，用来模拟生产者用户产生订单，并将产生的订单写入到单向 channel in 中
func produce(in chan<- OrderInfo) {
	for i := 0; i < 10; i++ {
		order := OrderInfo{id: i + 1}
		in <- order
	}
	close(in)
}

// 创建一个函数，用来模拟消费者，从单向 channel out 中读取订单信息
func consumer(out <-chan OrderInfo) {
	for order := range out {
		fmt.Printf("订单 id 为: %d\n", order.id)
	}
}

func main0601() {
	// 创建一个双向的 channel，注意定义的是一个结构体类型的 channel
	ch := make(chan OrderInfo)
	// 开启 go 程产生订单
	go produce(ch)
	// 调用 consumer 函数充当消费者获取订单信息
	consumer(ch)
}
