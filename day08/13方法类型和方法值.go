package main

import (
	"fmt"
)

type People1 struct {
	name string
	age  int
	sex  string
}

type Student3 struct {
	People1
	score int
}

func (p *People1) SayHello() {
	fmt.Printf("我叫 %s， 今年 %d 岁，我是 %s 生\n", p.name, p.age, p.sex)
}

func (s *Student3) SayHello(name string) {
	s.name = name
	fmt.Printf("我叫 %s， 今年 %d 岁，我是 %s 生, 我考试的分数为 %d 分\n", s.name, s.age, s.sex, s.score)
}

func main1301() {
	var stu Student3 = Student3{People1{"张三", 30, "男"}, 99}
	f1 := stu.SayHello
	f1("李四")
	fmt.Println(stu)
	//fmt.Printf("%T\n", f)
	/*f("李四")
	fmt.Println(stu)*/

	//var f1 hell
	//f1 = hello
	//f1()
}

//type hell func()
//
//func hello()  {
//	fmt.Println("hello")
//}
