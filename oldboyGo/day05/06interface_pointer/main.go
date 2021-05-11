package main

import "fmt"

// 使用值接收者和指针接收者的区别
type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// 使用值接收者实现了接口的所有方法
// func (c cat) move() {
// 	fmt.Println("猫在动！")
// }

// func (c cat) eat(food string) {
// 	fmt.Printf("%s在吃 %s\n", c.name, food)
// }

// 使用指针接收者实现接口的所有方法
func (c *cat) move() {
	fmt.Println("猫在动~~！")
}

func (c *cat) eat(food string) {
	fmt.Printf("%s在吃%s\n", c.name, food)
}

func main() {
	var a1 animal
	var c1 = cat{  // cat
		name: "tom",
		feet: 4,
	}
	var c2 = &cat{ // *cat
		name: "假老练",
		feet: 4,
	}
	a1 = &c1 // 实现 animal 这个接口的是 cat 的指针类型
	a1 = c2

	fmt.Println(a1)
}