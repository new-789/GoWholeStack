package main

import "fmt"

// 闭包
func f1(f func()) {
	fmt.Println("this is f1")
	f() // 此处等同于调用 f3 函数内部返回的匿名函数
}

func f2(x, y int) {
	fmt.Println("this is  f2")
	fmt.Println(x + y)
}

// 要求：f1(f2)
func f3(x, y int, f func(int, int)) func() {
	return func() {
		f(x, y) // 此处等同于调用 f2()
	}
}

func main() {
	ret := f3(100, 200, f2) // 将原来需要传递两个 int 类型参数包装成一个不需要传参的函数
	fmt.Printf("%T\n", ret)
	f1(ret)
}
