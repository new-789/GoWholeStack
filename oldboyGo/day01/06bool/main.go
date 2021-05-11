package main

import "fmt"

// 布尔值
func main() {
	b1 := true
	var b2 = false // 默认是 false
	fmt.Printf("%T\n", b1)
	fmt.Printf("%T value:%v\n", b2, b2)
}
