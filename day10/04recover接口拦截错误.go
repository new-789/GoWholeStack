package main

import "fmt"

func Demo(i int) {
	var arr [10]int

	// 通过匿名函数和 recover() 进行错误的拦截
	defer func() {
		// 可以从 panic 异常中重新获取控制权，程序还能继续执行
		//recover()

		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var p *int
	*p = 123

	// 如果传递超出数组下标的值会导致数组下标越界
	arr[i] = 100 // panic 系统处理，导致程序崩溃

	fmt.Println(arr)
}

func demo() {
	fmt.Println("Hello World")
}

func main0401() {
	Demo(11)
	demo()
}

// saskkasjh
