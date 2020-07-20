package main

import "fmt"

func main0601() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 定义数组的起始下标和结束下标
	start := 0
	end := len(arr) - 1

	for {
		// 判断如果起始下标比结束下标大，则说明交换完毕结束循环
		if start > end {
			break
		}
		// 做置换动作
		arr[start], arr[end] = arr[end], arr[start]
		// 使起始下标和结束下标不断向中间移动
		start++
		end--
	}
	fmt.Println("置换后的数组内容:", arr)
}
