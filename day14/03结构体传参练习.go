package main

import "fmt"

type Person2 struct {
	name      string
	age       int
	flg       bool
	intereset []string
}

// 通过函数参数初始化结构体
func initFunc(p *Person2) {
	p.name = "Nami"
	p.age = 18
	p.flg = true
	p.intereset = append(p.intereset, "购物", "跑步", "观影")
}

// 通过函数返回值初始化结构体
func initFunc1() *Person2 {
	p := new(Person2)
	p.name = "小强"
	p.age = 2
	p.flg = true
	p.intereset = append(p.intereset, "温度", "吓人", "墙角")
	return p
}

func main0301() {
	// 初始化结构体方式一调用函数代码
	/*
		var per Person2
		initFunc(&per)

		fmt.Println("per:", per)
	*/

	// 初始化结构体方式二调用函数代码
	p := initFunc1()
	fmt.Println("p:", p)
}
