package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(in chan<- int, idx int) {
	for i := 0; i < 50; i++ {
		num := rand.Intn(1000)
		fmt.Printf("%d th 生产者，生产：%d\n", idx, num)
		in <- num
	}
	close(in)
}

func consumer(out <-chan int, idx int) {
	for num := range out {
		fmt.Printf("==== %d th 消费者，消费 %d\n", idx, num)
	}
}

func main0601() {
	rand.Seed(time.Now().UnixNano())
	product := make(chan int)

	for i := 0; i < 5; i++ {
		go producer(product, i+1)
	}

	for i := 0; i < 5; i++ {
		go consumer(product, i+1)
	}

	for {

	}
}
