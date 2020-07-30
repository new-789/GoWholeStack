package main

import "fmt"

type Person struct {
	id    int
	name  string
	score int
}

// 定义一个函数，并设置一个结构体类型的形参
func test(s Person) {
	// 通过形参访问结构体中的内容
	fmt.Println(s.id)
	fmt.Println(s.name)
	fmt.Println(s.score)
}

func test1(s Person) {
	s.name = "玉皇大帝"
}

func main0601() {
	p := Person{101, "如来", 80}
	// 调用 test 函数并将 结构体 p 作为实参进行传递
	//test(p)
	test1(p)
	fmt.Println("查看在函数中是否修改成功:", p)
}

// 返回值方式修改结构体内容
func test2(s Person) Person {
	s.name = "玉皇大帝"
	return s
}

func main0602() {
	m := Person{101, "如来", 188}
	m = test2(m)
	fmt.Println(m)
}

func test3(s []Person) {
	// 通过切片的语法调用结构体中的成员进行修改内容
	s[0].name = "玉皇大帝"
	s[0].score = 80
	s[1].score = 60
}

func main0603() {
	// 定义结构体切片
	m := []Person{{101, "如来", 199}}
	m = append(m, Person{102, "齐天大圣", 90})
	test3(m)
	fmt.Println(m)
}
