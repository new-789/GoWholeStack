package main

import "fmt"

func main0301() {
	// 栈区 系统为每一个应用程序分配1M空间用来存储变量  在程序运行结束后系统会自动释放
	var s1 int = 10
	var s2 int = 20

	// 常量的定义，存储在内存的数据区  不用通过 & 运算符来访问
	const a int = 10

	fmt.Println(&s1)
	fmt.Println(&s2)

	//a = 20  // 常量的值不允许修改
	fmt.Println(a)
}

func main0302() {
	// 常量一般用大写字符表示
	const A int = 10
	b := 20
	c := A + b
	fmt.Println(c)
}

func main0303() {
	fmt.Println(123)
	fmt.Println("hello china")
	a := "hello"
	// "world" 为硬常量
	b := a + "world"
	fmt.Println(b)
}
