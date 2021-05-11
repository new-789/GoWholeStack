package main

import "fmt"

// 结构体占用一块连续的内存
type x struct {
	a int8 // int8 占用8个 bit 位 --> 1byte
	b int8
	c int8
}

func main() {
	m := x{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}
	fmt.Printf("%p\n", &(m.a))
	fmt.Printf("%p\n", &(m.b))
	fmt.Printf("%p\n", &(m.c))
}
