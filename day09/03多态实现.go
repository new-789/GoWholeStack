package main

import "fmt"

type people struct {
	name string
	age  int
	sex  string
}

type student struct {
	people
	score int
}

type teacher struct {
	people
	subject string
}

func (s *student) SayHello() {
	fmt.Printf("大家好，我叫 %s, 我是 %s 同学，今年 %d 岁， 我的考试成绩是 %d 分\n", s.name, s.sex, s.age, s.score)
}

func (t *teacher) SayHello() {
	fmt.Printf("各位同学好，我叫 %s ,我是 %s 老师， 今年 %d 岁，我主教的科目是 %s\n", t.name, t.sex, t.age, t.subject)
}

// 接口实现
type Peopleer interface {
	SayHello()
}

// 多态实现
// 多态是将接口作为函数参数  多态实现了接口的统一处理
func SayHello(p Peopleer) {
	p.SayHello()
}

func main0301() {
	var p Peopleer // 定义接口体类型的变量
	// 将结构体内存地址赋值给接口类型变量
	p = &student{people{"麻子", 20, "男"}, 88}
	SayHello(p) // 调用实现多态的函数，并将接口类型的变量作为实参进行传递

	p = &teacher{people{"傻X", 35, "女"}, "历史"}
	SayHello(p)
}
