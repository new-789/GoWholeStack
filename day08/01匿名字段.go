package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  string
}

// 结构体嵌套结构体
type Student struct {
	// 通过匿名字段实现继承
	Person // 结构体名称作为结构体的成员
	id     int
	score  int
}

func main0101() {
	// 定义结构体类型变量时直接结构体中的成员进行初始化，包括被嵌套的父类结构体
	var stu Student = Student{Person{"张三丰", 188, "男"}, 101, 100}
	//stu.id = 101
	// 结构体名称.父类成员信息
	//stu.name = "张三"
	//
	//stu.Person.name = "张三"
	//stu.score = 100
	//stu.sex = "男"
	//stu.age = 20

	fmt.Println(stu)
}
