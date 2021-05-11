package main

import (
	"fmt"
)

// 类型断言

func main() {
	var a interface{} // 定义一个空接口变量 a
	a = 100
	// 如何判断 a 保存的值的具体类型是什么
	// 类型断言
	// 1. x.(T)
	v, ok := a.(int8)
	if ok {
		fmt.Println("猜对了，a 的 int8 类型：", v)
	} else {
		fmt.Println("猜错了 a 不是 int8 类型")
	}

	// 2. switch
	switch vl := a.(type) {
	case int8:
		fmt.Println("a 是 in8 类型：", vl)
	case int16:
		fmt.Println("a 是 int16 类型：", vl)
	case int32:
		fmt.Println("a 是 int32 类型：", vl)
	case string:
		fmt.Println("string:", vl)
	case int:
		fmt.Println("int:", vl)
	default:
		fmt.Println("啥也不是")
	}
}
