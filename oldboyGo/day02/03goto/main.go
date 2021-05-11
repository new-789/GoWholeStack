package main

import (
	"fmt"
)

// goto
func main() {
	// 跳出多层 for 循环,常用
	var flag = false
	for i := 0; i < 10; i++ {
		for j := 'a'; j < 'z'; j++ {
			if j == 'c' {
				flag = true
				break // 跳出内层的 for 循环
			}
			fmt.Printf("%d-%c\n", i, j)
		}
		if flag {
			break // 跳出外层 for 循环
		}
	}
	fmt.Println("==============================")

	// goto 跳出多层循环语句
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				// 设置退出标签,调到指定的标签
				goto breakTag
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束 for 循环")
}
