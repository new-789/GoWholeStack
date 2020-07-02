package main

import "fmt"

func main0901() {
	// 逻辑非 !， 非真为假，非假为真，只能对 bool 类型的变量或者 bool 类型的表达式使用逻辑非运算符
	var a bool = false
	fmt.Println("逻辑运算符非", !a)
	a = true
	fmt.Println("逻辑运算符非", !a)
	b := 10
	c := 10
	fmt.Println("逻辑运算符非配合布尔表达式使用", !(b == c))

	//fmt.Println(!b)  // 此行报错
}

func main0902() {
	// 逻辑与 &&, 同真为真，其余为假
	a := 10
	b := 20
	c := 30
	d := 20
	fmt.Println("逻辑与 &&,两边都为真", a < b && c > d)
	fmt.Println("逻辑与 &&,其中一方为假", a < b && false) // 返回为假
	fmt.Println("逻辑与 &&， 两边都为假", a > b && c < d) // 返回为假
}

func main0903() {
	// 逻辑或 || ，同假为假，其余为真
	a := 10
	b := 20
	c := 30
	d := 20
	fmt.Println("逻辑与 ||,两边都为真：", a < b || c > d)
	fmt.Println("逻辑与 ||,一边都为真：", a < b || c < d)
	fmt.Println("逻辑与 ||,两边都为假：", a > b || c < d)
}

func main0904() {
	// 取地址运算符
	a := 10
	fmt.Println("取址运算符& :", &a)
	// 取值运算符 *
	p := &a // 将变量 a 指向的内存地址赋值给变量 p, p 称为指针变量
	// *p 取出变量 p 指向内存地址中存储的值
	fmt.Println("取值运算符 * ：", *p)
}
