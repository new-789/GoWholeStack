package main

import (
	"fmt"
	"runtime"
	"time"
)

func test() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 200)
		fmt.Println("===================匿名 go 程中的 test go 程")
	}
}

func main0001() {

	go func() {
		fmt.Println("--------------- 匿名 go 程111")
		go test()
		fmt.Println("--------------- 匿名 go 程222")
	}()

	for {
		runtime.GC()
	}
}
