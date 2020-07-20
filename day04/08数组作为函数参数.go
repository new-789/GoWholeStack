package main

import "fmt"

func test(arr [10]int) {
	for _, v := range arr {
		fmt.Println(v)
	}
	arr[2] = 123
	fmt.Printf("函数参数接收的数组地址:%p\n", &arr)
}
func main0801() {
	// 执行数组下标初始化元素，{8:5} 表示在下标为 8 的位置初始化元素为 5
	arr := [10]int{1, 2, 3, 4, 8: 5}
	// 数组作为函数参数时直接传递数组名即可
	test(arr)

	fmt.Println(arr)

	fmt.Printf("原数组地址:%p\n", &arr)
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

func main0802() {
	arr := [10]int{9, 1, 5, 6, 7, 3, 10, 2, 4, 8}
	// arr 数组变量为常量，但是它可以使用相同类型的数据对其进行重新赋值，
	// 此种赋值方式并不是单纯给 arr 这个变量本身赋值，而是对该数组变量控制的整体区间进行赋值
	arr = Bubble(arr)
	fmt.Println(arr)
}
