package main

import "fmt"

func main0401() {
	// 声明数组
	//var arr [10]int
	// 为数组下标为 0 和 1 的元素设置值
	//arr[0]=123
	//arr[1]=10

	// 声明数组并初始化元素内容
	var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 通过下标获取数组中的值
	//fmt.Println(arr[0])
	//fmt.Println(arr[1])
	//fmt.Println(arr[2])
	//fmt.Println(arr[3])
	//fmt.Println(arr)

	// 使用 for 循环遍历数组，该方法仅在玩吗知道数组长度时可用
	//for i:=0;i<10 ;i++  {
	//	fmt.Println(arr[i])
	//}

	// len() + for 循环遍历数组
	//for i:=0;i<len(arr) ;i++  {
	//	fmt.Println(arr[i])
	//}

	// 使用 range + for 遍历数组
	for i, v := range arr {
		fmt.Println("下标", i, ",元素值", v)
	}
}

func main0402() {
	// 定义数组时，可用初始化部分元素的值
	//var arr [10]int = [10]int{1,2,3,4,5}

	// 自动推导类型创建数组
	//arr := [10]int{1,2,3,4,5,6,7,8,9,10}

	// 使用自动推导类型，可用根据元素的个数创建数组
	arr := [...]int{1, 2, 3}
	for _, v := range arr {
		fmt.Println(v)
	}
	fmt.Printf("%T", arr)
}

func main0403() {
	//arr := [5]int{1,2,3,4,5}
	//arr[5] = 6 数组下标越界错误
	// arr[-1] = 9

	// 数组在定义后，元素个数就已经固定，不能添加

	// 数组是一个常量，不允许赋值操作，数组名代表整个数组
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%p\n", &arr)    // 直接取数组地址输出
	fmt.Printf("%p\n", &arr[0]) // 取数组中元素下标为 0 的地址输出
	fmt.Printf("%p\n", &arr[0]) // 取数组中元素下标为 1 的地址输出
}
