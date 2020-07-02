package main

import "fmt"

func main0301() {
	// 多重赋值
	a, b, c, d := 10, 20, 30, "hello"
	fmt.Println(a, b, c, d)
}

func main0302() {
	a, b := 10, 20
	// 交换变量的值
	a, b = b, a
	fmt.Println(a, b)
}
