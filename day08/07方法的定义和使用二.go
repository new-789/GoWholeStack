package main

import "fmt"

type stu struct {
	name string
	age  int
	sex  string
}

// 为 stu 结构体绑定一个方法
func (s stu) PrintInfo() {
	fmt.Println(s.name)
	fmt.Println(s.age)
	fmt.Println(s.sex)
}

func (s *stu) EditInfo(name, sex string, age int) {
	// 结构体指针直接调用结构体成员
	s.name = name
	s.age = age
	s.sex = sex
	//s.PrintInfo()  // 在一个方法中调用绑定的另一个方法
}

func main0702() {
	var s *stu   // 声明一个指针类型的结构体变量
	s = new(stu) // 为指针类型的结构体变量创建空间
	// 调用 EditInfo 方法并初始化字段的值
	s.EditInfo("薛宝钗", "女", 16)
	// 指针类型的结构体变量直接调用绑定方法中接收者不是指针类型的方法
	s.PrintInfo()
}

func main0701() {
	// 定义结构体类型变量，会将初始化的值自动传递给绑定的 printInfo 方法
	var s stu = stu{"布什", 48, "男"}
	// 对象.方法
	s.PrintInfo()
	// 结构体变量可以直接使用结构体指针对应的方法
	s.EditInfo("特朗普", "男", 59) // 等于 (&s).EditInfo("特朗普", "男", 59) 语法
	s.PrintInfo()
}
