package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
1. 使用 goroutine 和 channel 实现一个计算 int64 随机数个位数和的程序。
	1.1：开启一个 goroutine 循环生成 int64 类型的随机数，发送到 jobChan。
	1.2：开启 24 个 goroutine 从 jobChan 中取出随机数计算各位数的和，将结果发送到 resultChan。
	1.3：主 goroutine 从 resultChan 取出结果并打印到终端输出。
*/
// Job ...
type Job struct {
	value int64
}

// Result ...
type Result struct {
	Job *Job
	sum int64
}

var jobChan = make(chan *Job, 100)
var resultChan = make(chan *Result, 100)
var wg sync.WaitGroup

func generate(job chan<- *Job) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	// 循环生成 int64 类型的随机数
	for {
		x := rand.Int63()
		newJob := &Job{
			value: x,
		}
		job <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func calc(job <-chan *Job, result chan<- *Result) {
	defer wg.Done()
	// 从 jobChan 中取出随机数计算各位数的和，将结果发送到 resultChan
	for {
		v := <-job
		sum := int64(0)
		n := v.value
		// 取 int64 数的各位数
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &Result{
			Job: v,
			sum: sum,
		}
		result <- newResult
	}
}

func main() {
	wg.Add(1)
	go generate(jobChan)
	for i := 0; i <= 24; i++ {
		wg.Add(i)
		go calc(jobChan, resultChan)
	}

	for value := range resultChan {
		fmt.Printf("value:%d sum:%d\n", value.Job.value, value.sum)
	}
	wg.Wait()
}
