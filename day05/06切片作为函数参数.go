package main

import "fmt"

func test(s []int) {
	s[2] = 100
	//fmt.Printf("切片参数 s 的地址: %p\n", s)
}

func main0601() {
	s := []int{1, 2, 3, 4, 5}
	test(s)
	//fmt.Printf("主函数中切片 s 的地址: %p\n", s)
	fmt.Println(s)
}

// 冒泡排序
func BubbleSort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func main0602() {
	s := []int{9, 1, 5, 6, 7, 3, 10, 8, 2, 4}
	BubbleSort(s)
	fmt.Println(s)
}

// 函数中 append 方法注意示例
func test1(s []int) []int {
	s = append(s, 4, 5, 6)
	fmt.Printf("增加数据化地址：%p\n", s)
	return s
}

func main0603() {
	s := []int{1, 2, 3}
	fmt.Printf("原切片地址：%p\n", s)
	s = test1(s)
	fmt.Println(s)
	fmt.Printf("覆盖原切片后地址：%p\n", s)
}
