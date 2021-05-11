package main

import "fmt"

// 函数
// 函数：一段代码的封装
func f1() {
	fmt.Println("hello娜娜")
}

func f2(name string) {
	fmt.Println("hello", name)
}

// 带参数和返回值的函数
func f3(x int, y int) int {
	ret := x + y
	return ret
}

// 参数类型简写
func f4(x, y int) int {
	return x + y
}

// 可变参数
func f5(title string, y ...int) int {
	fmt.Println(y) // y 是一个 int 类型的切片
	return 1
}

// 命名返回值
func f6(x, y int) (sum int) {
	sum = x + y // 如果使用命名返回值，那么在函数中可以直接使用命名的返回值变量
	return      // return 后可以省略返回值变量
}

// go 语言中支持多个返回值
func f7(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func main() {
	f1()
	f2("娜娜")
	f2("流浪")
	fmt.Println(f3(100, 200))
	ret := f3(100, 200) // 调用函数
	fmt.Println(ret)
	f5("娜娜", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// 在一个命名的函数中不能在声明一个命名函数
}
