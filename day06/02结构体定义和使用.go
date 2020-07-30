package main

import "fmt"

// 在函数外部定义结构体, 作用域为全局
type Student struct {
	id   int
	name string
	sex  string
	age  int
	addr string
}

func main0201() {
	// 1. 通过结构体名定义结构体变量,定义的结构体变量包含结构体中的所有成员信息
	var s Student
	// 通过结构体变量访问结构体中的成员内容,并进行赋值
	s.id = 101
	s.name = "张飞"
	s.sex = "男"
	s.age = 88
	s.addr = "北京燕郊"
	fmt.Println("语法一: ", s)
	fmt.Println("语法一打印的字段内容: ", s.name) // 打印并输出结构体中某一个字段的内容

	// 定义结构体变量语法二，具体的字段名可不写，不写则必须和字段名一一对应
	var s1 Student = Student{102, "刘备", "男", 99, "荆州"}
	fmt.Println("语法二: ", s1)
	fmt.Println("语法二打印的字段内容: ", s1.name) // 打印并输出结构体中某一个字段的内容

	// 定义结构体变量语法三，具体的字段名可不写
	s2 := Student{
		addr: "东吴",
		id:   103,
		name: "小乔",
		age:  22,
		sex:  "女",
	}
	fmt.Println("语法三: ", s2)
	fmt.Println("语法三打印的字段内容: ", s2.name) // 打印并输出结构体中某一个字段的内容
}
