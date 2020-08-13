package main

import "fmt"

type People struct {
	name string
	age  int
	sex  string
}

type Student struct {
	People
	score int
}

type Teacher struct {
	People
	subject string
}

func (s *Student) SayHello() {
	fmt.Printf("大家好，我叫 %s, 我是 %s 同学，今年 %d 岁， 我的考试成绩是 %d 分\n", s.name, s.sex, s.age, s.score)
}

func (t *Teacher) SayHello() {
	fmt.Printf("各位同学好，我叫 %s ,我是 %s 老师， 今年 %d 岁，我主教的科目是 %s\n", t.name, t.sex, t.age, t.subject)
}

// 接口定义
// 接口定义了规则，方法实现了规则，接口是虚的，方法是实的
type Humaner interface {
	// 在接口中定义方法 函数的声明 没有具体实现，具体的实现根据对象的不同实现的方式也不同
	SayHello()
}

func main0201() {
	stu := Student{People{"张三", 18, "男"}, 89}
	teacher := Teacher{People{"张富国", 54, "男"}, "历史"}
	//stu.SayHello()
	//teacher.SayHello()
	var h Humaner // 定义接口类型变量
	h = &stu
	h.SayHello()

	h = &teacher
	h.SayHello()
}
