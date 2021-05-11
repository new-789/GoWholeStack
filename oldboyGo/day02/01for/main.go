package main

import "fmt"

// 流程控制之跳出 for 循环
func main() {
	// 当 i 等于 5 时跳出 for 循环
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break // 跳出 for 循环
	// 	}
	// 	fmt.Println(i)
	// }

	// 当 i 等于 5 时跳过此次循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // 跳过当前循环，继续下一次循环
		}
		fmt.Println(i)
	}
}
