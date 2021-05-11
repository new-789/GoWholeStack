package main

import "fmt"

// fmt 占位符
func main() {
	var n = 100
	// %T 查看类型
	fmt.Printf("%T\n", n)
	// %v 查看变量的值
	fmt.Printf("%v\n", n)
	// %d 查看十进制
	fmt.Printf("%d\n", n)
	// %b 查看二进制
	fmt.Printf("%b\n", n)
	// %o 查看八进制
	fmt.Printf("%o\n", n)
	// %x 查看十六进制
	fmt.Printf("%x\n", n)

	var s = "社会摇"
	// %s 查看字符串
	fmt.Printf("字符串: %s\n", s)
	fmt.Printf("%v\n", s)
	// %#v 查看变量的值，如果是字符串则会给结果自动添加 "" 号
	fmt.Printf("%#v\n", s)
}
