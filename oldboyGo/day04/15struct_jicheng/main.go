package main

import "fmt"

// 结构体继承：使用结构体模拟其他语言中的“继承”
type animal struct {
	name string
}

// 给 animal 实现一个动的方法
func (a animal) move() {
	fmt.Printf("%s 会动\n", a.name)
}

// 狗类
type dog struct {
	feet   uint8
	animal // 匿名结构体继承，animal 拥有的字段和方法，此时 dog 也同时拥有
}

// 给 dog 实现一个汪汪汪的方法
func (d dog) wang() {
	fmt.Printf("%s 再叫，汪汪汪...\n", d.name)
}

func main() {
	d1 := dog{
		animal: animal{name: "嘿嘿"},
		feet:   4,
	}
	fmt.Println(d1)
	d1.wang()
	d1.move() //
}
