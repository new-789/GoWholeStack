package main

import "fmt"

func main0601() {
	a := 10
	b := 3.14

	// 将不同类型转成相同的类型进行计算操作
	// float64(a) 将整型类型的数据转换为 float64 型
	c := float64(a) * b
	// int(b) 将浮点型的数据转换为整型数据，只保留浮点型的整数部分，舍弃小数部分，不会进行四舍五入
	d := a * int(b)
	fmt.Println(c)
	fmt.Println(d)
}

func main0602() {
	// 虽然 int32 和 Int64 都是整型，但是在 Go 语言中这两个类型的数据除了不能相互计算外，还不允许相互转换，如将 int32 换行为 int64 或将 int64 转换为 int32 是不被允许的
	var a int32 = 10
	var b int64 = 20
	c := int64(a) + b
	fmt.Println(c)
	fmt.Printf("%T", c)
}
