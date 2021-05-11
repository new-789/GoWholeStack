package main

import "fmt"

// 切片slice：切片的引用类型，都指向了底层的一个数组
func main() {
	// 定义切片
	var s1 []int    // 定义一个存放int 类型的元素切片
	var s2 []string //  定义一个存放string 类型的元素切片
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) // true
	fmt.Println(s2 == nil) // true

	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"北京", "上海", "广州"}
	fmt.Println(s1, s2)

	// 切片的长度和容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

	// 由数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s3 := a1[0:4] // 基于数组切割，左包含右不包含：左闭右开。[1 2 3 4]
	fmt.Println(s3)
	s4 := a1[1:6] // [2 3 4 5 6]
	fmt.Println(s4)
	s5 := a1[:4] // ==》[0:4]
	s6 := a1[3:] // ==> [3:len(a1)]
	s7 := a1[:]  // ==> [0:len(a1)]
	fmt.Println(s5, s6, s7)
	// 切片的容量是指底层数组的容量
	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5), cap(s5))
	// 底层数组从切片的第一个元素到最后一个元素的数量
	fmt.Printf("len(s6):%d cap(s6):%d\n", len(s6), cap(s6))

	// 切片再切片
	s8 := s6[5:]
	fmt.Printf("len(s8):%d cap(s8):%d\n", len(s8), cap(s8))
	a1[8] = 1000 // 修改了底层数组的值
	fmt.Printf("s6:%v\n", s6)
	fmt.Printf("s8:%v\n", s8)
}
