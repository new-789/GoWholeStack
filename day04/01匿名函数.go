package main

import "fmt"

func main0101() {
	//a := 10
	//b := 20

	// 匿名内部函数
	// f 是函数类型对应的变量
	//f := func (a, b int) {
	//	fmt.Println(a+b)
	//}
	//
	//f(a, b)
	//f(20, 30)

	f := func() {
		fmt.Println("hello world")
	}

	f()
	f()
	f()
}

func main0102() {
	a := 10
	b := 20
	// 带返回值的匿名函数
	v := func(a, b int) int {
		return a + b
	}(a, b)

	fmt.Println(v)
	fmt.Printf("%T", v)
}
