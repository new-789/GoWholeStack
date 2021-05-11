package main

import "fmt"

// if 条件判断
func main() {
	// age := 18
	// if age > 18 {
	// 	fmt.Println("澳门首家线上赌场开业了！")
	// } else {
	// 	fmt.Println("还未成年，回去写作业吧！")
	// }

	// 多个条件判断
	// if age > 35 {
	// 	fmt.Println("人到中年")
	// } else if age > 17 {
	// 	fmt.Println("青年")
	// } else {
	// 	fmt.Println("未成年")
	// }

	// 作用域：age 变量此时只在 if 条件判断语句中有效
	if age := 19; age > 18 {
		fmt.Println("澳门首家线上赌场开业了！")
	} else {
		fmt.Println("还未成年，回去写作业吧！")
	}
}
