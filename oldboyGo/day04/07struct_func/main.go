package main

import "fmt"

// 构造函数
type person struct {
	name string
	age  int
}

type dog struct {
	name string
}

// 构造函数：约定俗成使用  new 开头
// 构造函数返回的是结构体还是结构体指针
// 当结构体中的字段比较多时尽量使用结构体指针，减少程序的运行内存开销
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func main() {
	p1 := newPerson("龙马", 888)
	p2 := newPerson("悟净", 999)
	fmt.Println(p1, p2)

	d1 := newDog("旺财")
	fmt.Println(d1)
}
