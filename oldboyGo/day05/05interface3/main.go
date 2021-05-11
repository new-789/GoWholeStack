package main

import "fmt"

// 接口的实现
type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("猫在动！")
}

func (c cat) eat(food string) {
	fmt.Printf("%s在吃 %s\n", c.name, food)
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("鸡在动！")
}

func (c chicken) eat(food string) {
	fmt.Printf("鸡在吃 %s\n", food)
}

func main() {
	var a1 animal // 定义一个接口类型的变量
	 // 如果一个接口类型的变量中没有保存任何类型的值则它的值和类型为 nil
	fmt.Println(a1)
	fmt.Printf("%T\n", a1)
	var bc = cat{  // 定义一个 cat 类型的变量 bc
		name:"蓝猫",
		feet: 4,
	}
	a1 = bc
	a1.move()
	a1.eat("鱼")
	var kfc = chicken{
		feet: 2,
	}
	a1 = kfc
	a1.move()
	a1.eat("虫子")
}