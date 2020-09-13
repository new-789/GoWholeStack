package main

import "fmt"

/*
	死锁1：单 go 程自己死锁示例

*/
func main0101() {
	ch := make(chan int)
	ch <- 888
	num := <-ch
	fmt.Println("num =", num)
}

/*
	死锁2：go 程间 channel 访问顺序导致死锁示例
*/
func main0102() {
	ch := make(chan int)

	num := <-ch // 此处读 ch 管道产生阻塞下面的代码无法执行，所以在子 go 程中的写入操作无法执行
	fmt.Println("num =", num)

	go func() {
		ch <- 888
	}()
}

/*
	死锁3：多 go 程多 channel 交叉死锁示例
*/
func main0103() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 子 go 程中监听 ch1 并从中读出数据，然后将读出的数据写入到 ch2 中
	go func() {
		select {
		case num := <-ch1:
			ch2 <- num
		}
	}()

	// 主 go 程中监听 ch2 并从中读出数据，然后将读出的数据写入到 ch1 中
	for {
		select {
		case num := <-ch2:
			ch1 <- num
		}
	}
}
