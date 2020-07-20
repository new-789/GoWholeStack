package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main0901() {
	/*
		使用随机数的几个步骤：
			1. 导入头文件 math/rand 和 time 包
			2. 添加随机数种子,使每次获取到的随机数都不同
			3. 使用随机数
	*/
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子，用来保证每次产生的随机数不同；
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Intn(123)) // rand.Intn(123) 使用随机数，参数 123 用来设置随机数在 0~~122 之间的数字区间产生
	}
	fmt.Println(time.Now())
}

func BubbleSort(arr [10]int) [10]int {
	// 此处需要注意的是，将数组作为参数传递时，形参所接收的数组的个数必须与传递的实参保持一致
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main0902() {
	var arr [10]int // 声明一个 arr 数组

	rand.Seed(time.Now().UnixNano()) // 设置随机数种子
	for i := 0; i < len(arr); i++ {  // 通过循环将随机数添加至 arr 数组
		arr[i] = rand.Intn(100)
	}
	arr = BubbleSort(arr) // 调用 BubbleSort 方法对数组进行排序
	fmt.Println(arr)
}
