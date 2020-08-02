package main

import (
	"fmt"
)

type Student struct {
	id   int
	name string
	sex  string
	age  int
}

// 通过结构体指针间接操作结构体成员的信息
func main0902() {
	var stu Student = Student{101, "z张无忌", "男", 23}

	// 定义结构体指针
	var p *Student
	// 为结构体指针变量进行赋值
	p = &stu
	// 通过结构体指针变量间接操作结构体成员
	(*p).name = "赵敏"
	(*p).sex = "女"
	fmt.Println(stu)

	// 通过结构体指针变量直接操作结构体成员
	p.name = "张无忌"
	p.sex = "男"
	fmt.Println(stu)
}

// 结构体指针变量的定义
func main0901() {
	// 定义结构体变量
	// var 结构体名 结构体数据类型
	var stu Student = Student{101, "通天大圣", "不详", 1000}

	// 定义结构体指针指向变量的地址
	//var p *Student
	// 结构体指针变量与结构体建立关系，结构体指针变量指向结构体变量地址
	//p = &stu
	//fmt.Printf("结构体指针变量名类型：%T\n", p)

	fmt.Printf("结构体名地址:%p\n", &stu)
	fmt.Printf("结构体首元素地址:%p\n", &stu.id)
	p1 := &stu
	p2 := &stu.id // 不能使用该语法为结构体指针变量名赋值
	fmt.Printf("结构体名地址赋值给结构体指针变量名类型：%T\n", p1)
	fmt.Printf("结构体首个成员地址赋值给结构体指针变量名类型：%T\n", p2)
}

///
