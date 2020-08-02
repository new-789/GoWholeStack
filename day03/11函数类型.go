package main

import "fmt"

func test6() {
	fmt.Println("函数类型")
}

func test7(a, b int) {
	fmt.Println("带参数函数类型调用", a+b)
}

func test8() {
	fmt.Println("通过函数类型调用该函数")
}

func test9(a, b int) int {
	return a - b
}

// 关键字 type 可以定义函数类型,类型不同名称也不能相同，还可以为已存在类型起别名
// 使用 type 关键字定义函数类型
type FuncType func()

// 使用 type 关键字定义带参数的函数类型
type FuncTest func(int, int)

// 使用 type 关键字定义带参数与返回值的函数类型
type FuncReturn func(int, int) int

func main1101() {
	//test6()
	// 使用定义的函数类型定义变量
	var f FuncType
	f = test6 // 将一个函数名赋值给函数变量
	f()       // 通过函数变量调用 test6 函数
	//f = test7  // test7 函数类型与 FUNCTYPE 不一致不可以进行赋值
	f = test8
	f()

	var f1 FuncTest
	f1 = test7
	f1(10, 20)

	var f2 FuncReturn
	f2 = test9
	sub := f2(10, 20)
	fmt.Println("带参数与返回值的函数类型调用", sub)
}

func test10(a, b int) {
	fmt.Println("sum:", a+b)
}

func main1102() {
	// 此处注意，函数名不加括号不算调用函数
	fmt.Println(test10)
}
