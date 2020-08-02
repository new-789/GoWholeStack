package main

import "fmt"

// 定义一个 swap 函数,并接收两个指针int类型的参数
func swap(a, b *int) {
	// 通过取值运算符 * 交换两个参数的值
	*a, *b = *b, *a
}

func main0301() {
	// 定义两个int类型的变量，并进行赋值
	a := 10
	b := 20
	// 调用swap函数，并获取变量 a 和变量 b 的内存地址作为实参进行传递
	swap(&a, &b)
	// 查看在 swap 函数中交换了 a 和 b 之后，a 和 b 值的变化
	fmt.Println(a, b)
}

///
