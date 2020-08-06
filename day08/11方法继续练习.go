package main

import "fmt"

type Per struct { // 提取出来的父类信息，有公共属性
	name string
	age  int
	sex  string
}

// 定义程序猿结构体，用来存储自己独有的属性
type Dep struct {
	Per      // 匿名字段，继承父类结构体
	workTime int
}

// 定义记者结构体，用来存储自己独有的属性
type Rep struct {
	Per   // 匿名字段，继承父类结构体
	hobby string
}

func (p *Per) SayHello() {
	fmt.Printf("我是%s, 我的年龄是 %d, 我是 %s 生， ", p.name, p.age, p.sex)
}

func (r *Rep) SayHello() {
	// 当子类绑定了一个和父类同名的方法时，调用父类方法需使用 **对象.匿名字段.方法名** 语法
	r.Per.SayHello()
	fmt.Printf("我的爱好是 %s\n", r.hobby)
}

func (d *Dep) SayHello() {
	d.Per.SayHello()
	fmt.Printf("我的工作年限是 %d 年\n", d.workTime)
}

func main1101() {
	var r Rep = Rep{Per{"记者", 34, "男"}, "偷拍"}
	var d Dep = Dep{Per{"程序猿", 23, "男"}, 3}
	r.SayHello()
	d.SayHello()
}
