package main

import "fmt"

// 闭包练习2
func calc(base int) (func(int) int, func(int) int) {
	// base 是返回两个函数中的公共变量
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	add, sub := calc(10)
	fmt.Println(add(1), sub(2)) // 11, 9
	fmt.Println(add(3), sub(4)) // 12, 8
	fmt.Println(add(5), sub(6)) // 13, 7
}
