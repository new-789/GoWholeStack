package main

import "fmt"

// switch：简化大量的判断（一个变量和具体的值做比较）
func main() {
	var n = 5
	// if 写法
	// if n == 1 {
	// 	fmt.Println("大拇指")
	// } else if n == 2 {
	// 	fmt.Println("食指")
	// } else if n == 3 {
	// 	fmt.Println("中指")
	// }else if n == 4 {
	// 	fmt.Println("无名指")
	// }elseif n == 5 {
	// 	fmt.Println("小指")
	// }else {
	// 	fmt.Println("数据错误")
	// }

	// switch 简化写法
	switch n {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小指")
	default:
		fmt.Println("无效数字")
	}

	// switch 变种1
	switch n := 6; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇书")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println("。。。。。。")
	}

	// switch 变种2
	var age = 19
	switch {
	case age > 18:
		fmt.Println("成年人")
	case age < 19:
		fmt.Println("未成年")
	default:
		fmt.Println("好好学习")
	}

	// switch 中嵌套 fallthrough 很少用
	var s = "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough // 满足该分支后会继续执行下一个分支的内容，该关键字是为了兼容 c 语言，很少用
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println(".......")
	}
}
