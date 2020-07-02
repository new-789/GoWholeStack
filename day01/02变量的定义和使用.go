package main

import "fmt"

func main201() {
	// 变量的定义和使用
	// 声明 和 定义
	// 声明变量语法：var 变量名 数据类型, 变量在定义时存储不同的数据，需要不同的数据类型
	var a int // 声明变量,如果没有初始值，默认为 0
	// var b int = 10 // 定义变量
	a = 10
	a += 25
	fmt.Println(a)
}

func main0202() {
	// 极端圆的面积和周长
	//var PI float64 = 3.1415926
	//var r float64 = 2.5
	PI := 3.1415926
	r := 2.5
	// 面积 = PI * r * r
	//var s float64
	s := PI * r * r
	// 周长 = 2 * PI *r
	//var l float64
	l := 2 * PI * r
	fmt.Println("面积:", s)
	fmt.Println("周长:", l)
}

func main0203() {
	// 去市场买两斤黄瓜，价格为 5 元
	w := 2.0
	p := 5.0
	fmt.Println(w * p)
}

func main2004() {
	a := false // bool 布尔类型
	b := 10    // int 整型
	c := 3.14  // float64 浮点类型
	d := 'a'   // byte 字节
	e := "哈哈哈" // string 字符串
	fmt.Println("a=", a, "b=", b, "c=", c, "d=", d, "e=", e)
}
