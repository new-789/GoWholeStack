package main

import "fmt"

type DemoA struct {
	name string
	id   int
}

type DemoB struct {
	age int
	sex string
}

type DemoC struct {
	DemoA
	DemoB
	score int
}

func main0501() {
	var d DemoC = DemoC{DemoA{"戈洛文", 101}, DemoB{20, "男"}, 100}
	//d.name = "麻子"
	//d.id = 201
	//d.age = 17
	//d.sex = "男"
	//d.score = 98

	fmt.Println(d)
}
