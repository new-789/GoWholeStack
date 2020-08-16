package main

import (
	"fmt"
)

func main0301() {
	// defer 函数调用
	defer fmt.Println("Hello ")
	fmt.Println("World ")
}

func test4() {
	fmt.Println("Hello World")
}

var value int

func test5(a, b int) {
	value = a + b
}

func test6(value int) {
	fmt.Println(value)
}

func main0302() {
	test4()
	// 函数中如果有返回值不能使用 defer 调用
	defer test5(10, 20)
	test6(value)
}

// defer 与匿名哈数结合使用示例
func main0303() {
	a := 10
	b := 20
	// 没有参数，主函数结束前执行匿名函数，会去找主函数中 a 和 b 两个变量,所以该匿名函数的结果为 100 和 200
	//defer func () {
	//	fmt.Println(a)
	//	fmt.Println(b)
	//}()

	// 有参数，主函数结束前执行匿名函数，由于在存入栈区前将 a 和 b 两个变量传递给了匿名函数,所以该匿名函数的结果为 10 和 20
	defer func(a, b int) {
		fmt.Println(a)
		fmt.Println(b)
	}(a, b)

	a = 100
	b = 200

	fmt.Println(a)
	fmt.Println(b)
}
