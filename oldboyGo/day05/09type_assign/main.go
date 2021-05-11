package main

import "fmt"

// 类型断言1
func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string) // 类型直接断言
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Println("传进来的是一个字符串", str)
	}
}

// 类型断言2
func assign2 (a interface{}) {
	fmt.Printf("%T\n", a)
	// switch 做断言判断语法
	switch v := a.(type) {
	case string:
		fmt.Println("传进来的是一个字符串", v)
	case int:
		fmt.Println("传进来的是一个 int", v)
	case int64:
		fmt.Println("传进来的是一个 int64", v)
	case bool:
		fmt.Println("传进来的是一个 bool：", v)
	default:
		fmt.Println("以上类型都不对，不知道传了个啥")
	}
}

func main() {
	assign(100)
	assign2(888)
	assign2(int64(999))
	assign2("哈哈哈")
	assign2(true)
}