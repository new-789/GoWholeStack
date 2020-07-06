package main

import "fmt"

func main1201() {
	a := 10
	// 匿名内部函数,也有一个作用域
	{
		a := 20
		fmt.Println(a)
	}
	fmt.Println(a)
}

func main() {
	demo2()
	//a := 234
	fmt.Println(a)
}

// 全部变量 a, 在函数外部定义的变量
var a int

func demo2() {

	{
		a = 123
	}
	fmt.Println(a)
}
