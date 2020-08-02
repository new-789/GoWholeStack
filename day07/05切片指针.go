package main

import "fmt"

func main0503() {
	// 定义切片
	var slice []int = []int{1, 2, 3, 4, 5}
	// 将切片的内存地址赋值给指针
	var p *[]int = &slice
	fmt.Printf("原切片地址: %p\n", slice)
	// 通过指针往切片中增加内容
	*p = append(*p, 6, 7, 8, 9, 10)
	fmt.Println("增加内容后切片中的内容：", slice)
	fmt.Printf("增加内容后切片的地址: %p\n", slice)
}

func main0502() {
	var slice []int = []int{1, 2, 3, 4, 5}
	fmt.Println("原切片：", slice)
	var p *[]int = &slice
	fmt.Println("通过指针变量打印切片中元素: ", (*p)[1])
	// 通过指针修改切片中元素下标为 1 的值
	(*p)[1] = 100
	fmt.Println("指针修改切片内容后的切片：", slice)
}

func main0501() {
	var slice []int = []int{1, 2, 3, 4, 5}
	var p *[]int = &slice // 将切片内存地址赋值给指针变量
	fmt.Printf("指针变量内存地址: %p\n", p)
	fmt.Printf("切片变量内存地址: %p\n", slice)
}

///
