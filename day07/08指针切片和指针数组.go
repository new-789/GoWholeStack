package main

import "fmt"

//指针切片
func main0802() {
	// 定义指针切片
	var slice []*int
	a := 10
	b := 20
	c := 30
	d := 40
	// 往指针切片中添加数据
	slice = append(slice, &a, &b, &c, &d)
	fmt.Println(slice) // 打印指针切片存错的所有内容
	// 通过 for 循环遍历指针切片语法一
	fmt.Println("--------- 普通 for 循环语法 ----------")
	for i := 0; i < len(slice); i++ {
		fmt.Printf("每一个数据的类型: %T\n", slice[i])
		fmt.Printf("每一个具体的值：%d\n", *slice[i])
	}
	fmt.Println("------- for range 语法----------")
	// 通过 for 循环遍历指针切片语法二
	for i, v := range slice {
		fmt.Printf("下标: %d, 具体的值：%d\n", i, *v)
	}
}

// 指针数组
func main0801() {
	//var arr [3]int  定义数组变量
	// 定义指针数组变量
	var arr [3]*int
	a := 10
	b := 20
	c := 30
	// 往指针数组中增加数据，注意往指针数组中存储的内容必须是一个地址类型的数据
	arr[0] = &a
	arr[1] = &b
	arr[2] = &c

	fmt.Println(arr)
	// 通过指针变量修改变量 b 的值
	*arr[1] = 200
	fmt.Println(b)
	// 循环打印指针数组中保存的每一个指针变量类型具体的值
	for i := 0; i < len(arr); i++ {
		// *arr[i] 获取指针数组中每一个元素具体的值语法
		fmt.Println("下标:", i, "具体的值：", *arr[i])
	}
}

///
