package main

import (
	"fmt"
	"unsafe"
)

type Person1 struct {
	name string
	age  int
	sex  byte
}

func main0201() {
	// 创建结构体指针变量方式一
	/*
		var p1 *Person1 = &Person1{"n1", 20, 'f'}
		fmt.Println("p1", p1)
	*/

	// 创建结构体指针变量方式二
	/*
		var tmp Person1 = Person1{"n1", 20, 'f'}
		var p2 *Person1
		p2 = &tmp
		fmt.Println("p2", p2)
	*/

	// 创建结构体指针变量方式三
	p3 := new(Person1)
	p3.name = "A1"
	p3.age = 22
	p3.sex = 'm'
	fmt.Println("src p3 = ", p3)
	fmt.Println("main p3 size = ", unsafe.Sizeof(p3))

	// 指针结构体地址
	/*
		fmt.Printf("p3 = %p\n", p3)  // 打印输出 p3 的值
		fmt.Printf("&p3.name = %p\n", &p3.name)  // 打印输出结构体首元素地址
	*/

	// 指针结构体传参
	// 由于 new 方式创建出来的结构体变量为指针类型的结构体变量，所以我们直接将其传递给接收指针类型的形参即可
	test1(p3)
	fmt.Println("func_after p3 = ", p3)
}

func test1(p *Person1) {
	fmt.Println("test2 p size = ", unsafe.Sizeof(p))
	p.name = "luffy"
	p.age = 999
	p.sex = 'f'
}
