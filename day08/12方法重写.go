package main

import "fmt"

type per struct {
	name string
	age  int
	sex  string
}

type teacher struct {
	per
	subject string
}

type student1 struct {
	per
	score int
}

func (p *per) PrintInfo() {
	fmt.Printf("我叫%s, 今年 %d 岁了，性别 %s\n", p.name, p.age, p.sex)
}

// 子类重写父类方法
func (t *teacher) PrintInfo() {
	fmt.Printf("我叫 %s, 今年 %d 岁了, 我是%s老师，主教的科目是 %s\n", t.name, t.age, t.sex, t.subject)
}

func (s *student1) PrintInfo() {
	fmt.Printf("我叫 %s, 今年 %d 岁了， 我是 %s 生，我考试成绩为: %d\n", s.name, s.age, s.sex, s.score)
}

func main1201() {
	var teacher teacher = teacher{per{"郑成功", 40, "男"}, "航海"}
	var stu student1 = student1{per{"小明", 19, "男"}, 100}
	// 调用子类方法
	teacher.PrintInfo()
	stu.PrintInfo()
	// 调用父类方法
	teacher.per.PrintInfo()
	stu.per.PrintInfo()
}
