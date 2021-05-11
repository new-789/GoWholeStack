package main

import "fmt"

// 浮点数
func main() {
	// math.MaxFloat32 // float32 最大值
	f1 := 1.23456
	fmt.Printf("%T\n", f1) // 默认 go 语言中的小数都是 float64 类型
	f2 := float32(1.23456) // 显示声明 float32 类型
	fmt.Printf("%T\n", f2)
	// f1 = f2 // float32 类型的值不能直接赋值给 float64 类型
}
