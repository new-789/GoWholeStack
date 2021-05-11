package main

import (
	"fmt"
	"time"
)

//

func work(id int, jobs <-chan int, result chan<- int) {
	for i := range jobs {
		fmt.Printf("worker :%d start job: %d\n", id, i)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, i)
		result <- i * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)
	// 开启三个 goroutine
	for i := 0; i <= 3; i++ {
		go work(i, jobs, result)
	}
	// 5个任务
	for j := 0; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	// 输出结果
	for k := 0; k <= 5; k++ {
		<-result
	}
}
