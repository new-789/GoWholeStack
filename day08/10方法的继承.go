package main

import "fmt"

type Person1 struct {
	id   int
	name string
	age  int
	sex  string
}

type Student2 struct {
	//p     Person1 // 实名字段，结构体变量  结构体类型
	Person1
	score int
}

// 帮顶给不同结构体的方法可以同名
// 定义方法绑定给 Peron1 结构体
func (p *Person1) SayHello() {
	fmt.Println("大家好，我是：", p.name, "我今年：", p.age, "我是:", p.sex)
}

// 定义方法绑定给 Student2 结构体
func (s *Student2) SayHello() {
	fmt.Println("大家好，我是：", s.name, "我今年：", s.age, "我是:", s.sex)
}

func main1001() {
	var stu Student2 = Student2{Person1{203, "贾宝玉", 18, "男"}, 99}
	var p Person1 = Person1{102, "取名好难", 20, "男"}
	p.SayHello()
	//stu.p.name = "贾宝玉"
	//stu.p.age = 18
	//stu.p.sex = "男"
	//stu.p.id = 203
	//stu.score = 100

	// 子类结构体继承父类结构体，允许使用父类结构体成员，还允许使用父类的方法
	stu.SayHello()

	//fmt.Println(stu)
}
