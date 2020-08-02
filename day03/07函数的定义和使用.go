package main

import "fmt"

func add0701(a int, b int) {
	sum := a + b
	fmt.Println(sum)
}

func swap(a, b int) {
	a, b = b, a // 在此处交换 a 和 b 的值
	fmt.Println("SWAP 函数中的结果", a, b)
}

func main0702() {
	a1 := 10
	b1 := 20
	//add(a1, b1)
	swap(a1, b1)
	fmt.Println("入口函数中的结果", a1, b1)
}
