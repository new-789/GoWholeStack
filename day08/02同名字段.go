package main

import "fmt"

type person struct {
	name string
	age  int
	sex  string
}

type student struct {
	person
	id    int
	name  string // 与 person 同名了
	score int
}

func main0201() {
	// 定义结构体类型变量时，直接初始化自身结构体中的字段，以及初始化内嵌结构体(父类)字段
	var stu student = student{person{"张三疯", 188, "男"}, 101, "张无忌", 99}
	// 采用就近原则，使用的是子类结构体中的字段名
	//stu.name = "张三丰"
	// 初始化内嵌结构体字段语法
	//stu.person.name = "张三疯"
	//
	//stu.id = 101
	//stu.score = 99
	// 初始化内嵌结构体字段语法
	//stu.person.age = 18
	//stu.person.sex = "男"

	fmt.Println(stu)
}
