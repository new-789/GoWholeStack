package main

import "fmt"

func main0501() {
	arr := [10]int{2, 5, 9, 6, 10, 4, 7, 3, 8, 1}
	max := arr[0] // 假设最重的那只小猪为数组中的第一个元素
	for i := 1; i < len(arr); i++ {
		// 通过循环不断和数组中其它元素做对比，只要遇见比 max 值大的则对 max 重新赋值，最后求出体重最重的小猪
		if max < arr[i] {
			max = arr[i]
		}
	}
	fmt.Println("最重的小猪为:", max)
}
