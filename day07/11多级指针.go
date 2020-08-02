package main

import "fmt"

func main1101() {
	a := 10
	p1 := &a  // 一级指针
	p2 := &p1 // 二级指针，用来存储一级指针的内存地址
	p3 := &p2 // 三级指针，用来存储二级指针的内存地址

	//fmt.Printf("变量 a 的类型：%T\n", a)
	//	//fmt.Printf("一级指针变量 p 的类型：%T\n", p)
	//	//fmt.Printf("二级指针变量 pp 的类型：%T\n", pp)
	//	//fmt.Printf("三级指针变量 pp 的类型：%T\n", ppp)

	// 通过三级指针获取变量 a 的值
	fmt.Println(***p3)
	// 通过二级指针获取变量 a 的值
	fmt.Println(**p2)
	// 通过一级指针获取变量 a 的值
	fmt.Println(*p1)

	// 通过三级指针修改变量 a 的值
	***p3 = 100
	fmt.Println("修改之后的值：", a)
}

///
