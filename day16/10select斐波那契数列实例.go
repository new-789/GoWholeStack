package main

import (
	"fmt"
	"runtime"
)

// 创建一个 斐波那契函数，用来打印输出斐波那契数列
func Fibonacci(ch <-chan int, quit <-chan bool) {
	for {
		select {
		case num := <-ch:
			fmt.Print(num, " ")
		case <-quit:
			runtime.Goexit() // 退出子 go 程,等效于 return
		}
	}
}

func main1001() {
	ch := make(chan int)
	quit := make(chan bool)

	go Fibonacci(ch, quit)

	x, y := 1, 1 // 定义斐波那契数列的两个初始值，x 相对于前一个数，y 相当于后一个数
	for i := 0; i < 20; i++ {
		/*
			将 x 的值添加到队列，因为 x 永远是前一个值，所以得到的结果即为斐波那契数列，
			如果写入 y 则刚开始会少一个1 结果不同，虽然也是按照斐波那契计算，但刚开始就不对了就不算是斐波那契数列了
		*/
		ch <- x
		x, y = y, x+y // 计算斐波那契数列，将 y 的值赋值给 x,y 的值更新为 x 加 y 的值,如此循环结果即为前一个数加上后一个数
	}
	quit <- true
}
