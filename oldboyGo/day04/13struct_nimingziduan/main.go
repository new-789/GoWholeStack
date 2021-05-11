package main

import "fmt"

// 结构体匿名字段：适用场景，字段比较少且简单的场景，不常用
type person struct {
	string
	int
}

func main() {
	p1 := person{
		"悟空",
		88,
	}
	fmt.Println(p1)
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}
