package main

import "fmt"

// 嵌套结构体
type address struct {
	province string
	city     string
}

type workPlace struct {
	province string
	city     string
}

type perosn struct {
	name string
	age  int
	// addr address 嵌套结构体
	address   // 匿名嵌套结构体1
	workPlace // 匿名嵌套结构体2
}

type company struct {
	name string
	address
}

func main() {
	p1 := perosn{
		name: "悟空",
		age:  9000,
		address: address{ // 初始化匿名嵌套结构体方法
			province: "东胜神州",
			city:     "花果山",
		},
		workPlace: workPlace{
			province: "天庭",
			city:     "馬圈",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.name)
	// fmt.Println(p1.address.city) // 获取嵌套结构体中字段的值

	// 通过匿名嵌套结构体可直接像访问自己结构体中自当方法获取嵌套结构体中字段的值
	// fmt.Println(p1.city) // 查询顺序，先从自己结构体中查找字段，如找不到则去匿名嵌套结构体中查抄

	// 如一个结构体中同时嵌套了两个匿名结构体，且两个嵌套结构体中都用相同的字段
	// 那么在访问时则需要指定具体访问那个嵌套结构体中的字段
	fmt.Println(p1.address.city)
	fmt.Println(p1.workPlace.city)
}
