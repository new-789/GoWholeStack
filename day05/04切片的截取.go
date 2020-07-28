package main

import "fmt"

func main0401() {
	s := []int{1, 2, 3, 4, 5, 6, 7}

	//slice := s[0:]  截取从下表 0 开始到最后一个下标结束切片之间的数据
	//slice := s[:3]  截取从切片中第一个下标开始至下标为 3 结束之间的数据，不包含下标为3的数据
	//slice := s[2:5]   截取切片中下标为 2 至下标为 5 之间的数据，不包含下标为 5 的数据
	slice := s[0:2:5]
	fmt.Println(slice)
}

func main0402() {
	s := []int{1, 2, 3, 4, 5}
	slice := s[2:5]
	slice[1] = 123

	fmt.Println("切片后", slice)
	fmt.Println("原切片", s)

	fmt.Printf("原切片地址:%p\n", s)
	fmt.Printf("新切片地址:%p\n", slice)
}

func main0403() {
	s := []int{1, 2, 3, 4, 5}

	slice := s[:]
	fmt.Println(slice)
}
