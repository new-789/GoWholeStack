package main

import "fmt"

// 引出接口的实例
type speaker interface { // 定义一个能叫的类型
	speak() // 只要实现了 speak 方法的变量，都是 speaker 类型，方法签名
}

type cat struct{}

type dog struct{}

type person struct {}

func (c cat) speak() {
	fmt.Println("喵喵喵")
}

func (d dog) speak() {
	fmt.Println("汪汪汪")
}

func (p person) speak() {
	fmt.Println("啊啊啊~")
}

func da(x speaker) {
	// 接收一个参数，传进来谁我就打什么
	x.speak() // 挨打了就要叫
}

// 再编程中会遇到的问题：我们不关心一个变量是什么类型，只关心他能调用它的什么方法
func main() {
	var c1 cat
	var d1 dog
	var p1 person
	da(d1)
	da(c1)
	da(p1)

	var ss speaker // 定义一个接口类型：speaker 的变量 ss
	ss = c1
	ss = d1
	ss = p1
	fmt.Println(ss)
}
