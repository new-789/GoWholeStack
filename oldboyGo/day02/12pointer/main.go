package main

import "fmt"

// 指针
func main() {
	// & 取地址
	// n := 10
	// p := &n
	// fmt.Printf("%T\n", p) // *int ：int 类型的指针
	// // * 根据地址取值
	// m := *p // 根据地址取值
	// fmt.Println(m)
	// fmt.Printf("%T\n", m)

	var a1 *int
	fmt.Println(a1)
	var a2 = new(int) // new 函数申请一个内存地址
	fmt.Println(a2)
	fmt.Println(*a2)
	*a2 = 100
	fmt.Println(*a2)
}
