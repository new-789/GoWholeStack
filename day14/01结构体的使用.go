package main

import (
	"fmt"
	"unsafe"
)

type Person struct { // 定义一个结构体类型
	name string
	sex  byte
	age  int
}

type Student struct {
	Person
	id    int
	score int
}

func main0101() {
	// 1. 顺序初始化
	/*
		var man Person = Person{"Andy", 'm', 18}
		fmt.Println("man", man)
	*/

	// 2. 指定成员初始化
	/*
		girl := Person{age: 20, sex: 'f'}
		fmt.Println("girl:", girl)
		// 索引成员变量 "."
		fmt.Printf("girl.name = %q\n", girl.name)
	*/

	// 结构体中普通变量的赋值和使用
	/*
		var girlAndMan Person
		girlAndMan.name = "不男不女"
		girlAndMan.sex = '?'
		girlAndMan.age = 26
		fmt.Printf("girlAndMan = %v\n", girlAndMan)
		girlAndMan.age = 22  // 通过结构体变量调用结构体成员修改并修改信息值
		fmt.Printf("girlAndMan = %v\n", girlAndMan)
	*/

	// 结构体变量比较
	/*
		var p1 Person = Person{"Andy", 'm', 18}
		var p2 Person = Person{"Andy", 'm', 18}
		var p3 Person = Person{"Andy", 'm', 118}
		// 比较两个结构体变量，并打印输出
		fmt.Println("p1 == p2 ? -->", p1 == p2)
		fmt.Println("p2 == p3 ? -->", p2 == p3)
	*/

	// 相同类型结构体赋值,即将一个结构体变量数据通过 = 赋值给另一个结构体变量，前提是结构体类型必须一致
	/*
		var tmp Person
		fmt.Println("tmp", tmp)
		tmp = p3
		fmt.Println("tmp", tmp)
	*/

	// 函数内部使用结构体传参
	var temp Person
	// unsafe.Sizeof() 方法可以求一个变量的大小
	fmt.Println("man temp size:", unsafe.Sizeof(temp))
	test(temp)
	fmt.Println("src temp", temp)

	// 结构体地址
	fmt.Printf("&temp = %p\n", &temp)
	fmt.Printf("&temp.name = %p\n", &temp.name)

	fmt.Println("bool size:", unsafe.Sizeof(true))

}

func test(man Person) {
	fmt.Println("test man size:", unsafe.Sizeof(man))
	man.name = "Rose"
	man.age = 22
	fmt.Println("test man:", man)
}
