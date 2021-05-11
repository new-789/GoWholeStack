package main

import "fmt"

//
func main() {
	var a int
	a = 100
	b := &a
	fmt.Printf("type a: %T type b: %T\n", a, b)
	// 将变量 a 十六进制的内存地址打印出来
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", b) // b 的值
	fmt.Printf("%v\n", b)
	fmt.Printf("%p\n", &b) // b 的内存地址
}
