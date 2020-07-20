package main

import "fmt"

func main0701() {
	arr := [10]int{4, 2, 8, 10, 0, 5, 7, 1, 3, 9}
	for i := 0; i < len(arr)-1; i++ { // 外层控制行
		// 内层控制列
		for j := 0; j < len(arr)-1-i; j++ {
			// 比较相邻两个元素，满足条件交换数据，升序使用大于号，降序使用小于号
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("排序完成后结果：", arr)
}
