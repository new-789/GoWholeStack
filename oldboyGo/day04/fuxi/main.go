package main

import "fmt"

func test(name string) {
	fmt.Println("hello,", name)
}

// 函数作为参数
func lixiang(f func(string), name string) {
	f(name)
}

// 函数作为返回值
func test2() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func low(f func()) {
	f()
}

// 闭包
func bi(f func(string), name string) func() {
	return func() {
		f(name)
	}
}

func main() {
	lixiang(test, "联想")
	ret := test2()
	fmt.Printf("%T\n", ret)
	sum := ret(100, 200)
	fmt.Println(sum)

	// 闭包演示
	fc := bi(test, "元帅")
	low(fc)
}
