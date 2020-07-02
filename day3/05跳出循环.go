package main

import "fmt"

func main0501() {
	// break 的使用
	for i := 1; i < 5; i++ {
		fmt.Println(i)
		if i == 3 {
			break
		}
	}
}

func main0502() {
	for i := 1; i <= 100; i++ {
		if i%2 == 1 { // 用 2 取模等于 1 表示为奇数，则跳出本次循环，进入下一次循环
			continue
		}
		fmt.Println(i)
	}
}
