package main

import "fmt"

// 结构体
type person struct { // 创建结构体类型语法
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var p person // 通过结构体类型定义变量 p
	// 通过字段赋值
	p.name = "理想"
	p.age = 8888
	p.gender = "男"
	p.hobby = []string{"篮球", "足球", "双色球"}
	fmt.Println(p)
	// 访问变量 p 的字段
	fmt.Printf("%T\n", p)
	fmt.Println(p.name)

	var p2 person
	p2.name = "娜娜"
	p2.age = 18
	fmt.Printf("types: %T value: %v\n", p2, p2)

	// 匿名结构体：多用于临时场景
	var s struct {
		name string
		age  int
	}
	s.name = "哈哈"
	s.age = 100
	fmt.Printf("type: %T value: %v\n", s, s)
}
