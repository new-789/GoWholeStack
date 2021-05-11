package main

import "fmt"

// 同一个结构体可以实现多个接口
// 接口还可以嵌套
type animal interface {
	mover // animal 接口嵌套mover接口
	eater // animal 接口嵌套 eater 接口
}

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// cat 实现了 move 接口
func (c *cat) move() {
	fmt.Println("猫在动")
}

// cat 实现了 eat 方法
func (c *cat) eat(food string) {
	fmt.Printf("%s在吃%s\n", c.name, food)
}

func main() {
	var a1 animal // 定义 animal 接口类型变量
	var c1 = &cat{"tom", 4,}
	a1 = c1
	a1.move()
	a1.eat("咸鱼")
}