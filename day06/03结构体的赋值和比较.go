package main

import "fmt"

type Students struct {
	id   int
	name string
	sex  string
	age  int
	addr string
}

func main0301() {
	s := Students{101, "貂蝉", "女", 20, "山西"}

	// 结构体变量赋值
	s1 := s
	// 通过接收赋值的变量修改结构体中元素为 age 的值
	s1.age = 18
	fmt.Println("赋值后的变量s1: ", s1)
	fmt.Println("原结构体变量s: ", s)
}

func main0302() {
	s := Students{101, "貂蝉", "女", 20, "山西"}

	s1 := s
	s1.name = "王昭君"
	// 比较两个结构体成员内容是否一致
	if s == s1 {
		fmt.Println("相同的两个结构体")
	} else {
		fmt.Println("不相同的两个结构体")
	}

	// 比较两个结构体中成员内容的大小关系，前提是存储的数据类型支持大于或者小于的关系运算符比较
	if s.age > s1.age {
		fmt.Println("s中的 age 大于 s1 中的 age")
	} else {
		fmt.Println("s中的 age 小于 s1 中的 age")
	}
}
