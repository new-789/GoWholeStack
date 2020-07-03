package main

import "fmt"

//func sub(a, b int) int {
//	sum := a - b
//	// return 表示函数的结束，它后面的代码不会执行，同时 return 也会将函数的返回值传递给主函数
//	return sum
//}

//func sub(a, b int) int  {
//	// return 后面可直接写一个表达式，将表达式的结果作为返回值
//	return a - b
//}

func sub(a, b int) (sum int) {
	sum = a - b
	return
}

func main1001() {
	value := sub(10, 20)
	fmt.Println(value)
}

func test5() (a int, b int, c int) {
	a, b, c = 1, 2, 3
	return
}

func main1002() {
	// 如果函数没有参数，在函数调用时 () 必须写
	// 函数的返回值可以为主调函数进行赋值操作， 可以通过返回值改变实参数据
	// 如果函数有多个返回值， 但是不使用其中的某一个返回值，可以用匿名变量接收数据
	a, _, c := test5()
	fmt.Println(a)
	//fmt.Println(b)
	fmt.Println(c)
}
