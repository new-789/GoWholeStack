package main

import "fmt"

// 整型
func main() {
	// 十进制
	var i1 = 100
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1) // 将十进制转换成二进制
	fmt.Printf("%o\n", i1) // 将十进制转换成八进制
	fmt.Printf("%x\n", i1) // 将十进制转化成十六进制
	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)

	// 十六进制数
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)

	// 查看变量类型
	fmt.Printf("%T\n", i3)

	// 声明 int8 类型的变量
	i4 := int8(9) // 明确指定 Int8 类型，否则默认位 int 类型
	fmt.Printf("%T\n", i4)
}
