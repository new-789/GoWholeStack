package main

import "fmt"

// 接口示例2
// 不管是什么品牌的车，都能跑
// 定义一个car接口类型，不管是什么结构体，只要实现了 run 方法，则都是 car 类型
type car interface {
	run()
}

type falali struct {
	brand string
}

type baoshijie struct {
	brand string
}

func (f falali) run() {
	fmt.Printf("%s 速度 80迈\n", f.brand)
}

func (b baoshijie) run() {
	fmt.Printf("%s 速度 180 迈\n", b.brand)
}

// drive 函数接收一个 car 类型的变量
func drive(c car) {
	c.run()
}

func main() {
	var f1 = falali{
		brand: "法拉利",
	}	
	var b1 = baoshijie{
		brand: "保时捷",
	}

	drive(f1)
	drive(b1)
}