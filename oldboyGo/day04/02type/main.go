package main

import "fmt"

// 自定义类型和类型别名

// type 后面跟的是类型

type MyInt int     // 创建自定义类型
type yourInt = int // 定义类型别名

func main() {
	var n MyInt

	n = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n) // main.MyInt

	var m yourInt // 通过类型别名定义变量
	m = 100
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	var c rune
	c = '国'
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
