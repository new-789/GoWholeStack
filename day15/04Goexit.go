package main

import (
	"fmt"
	"runtime"
)

func test() {
	// 在 Goexit 之前注册的 defer 在结束当前 go 程后会被执行
	defer fmt.Println("ccccccccccccccc")
	runtime.Goexit() // 退出当前 go 程
	//return
	fmt.Println("ddddddddddddddd") // 不会执行
}

func main0401() {

	go func() {
		// 在 Goexit 之前注册的 defer 在结束当前 go 程后会被执行
		defer fmt.Println("aaaaaaaaaaaaaaaaa")
		test()
		fmt.Println("bbbbbbbbbbbbbbbb") // 不会执行
	}()

	for {

	}
}
