package main

import "fmt"

func main0401() {
	const (
		a = iota // 0 表示静止
		b = iota // 1 表示移动
		c = iota // 2 表示普通攻击
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// 定义一个变量表示游戏中人物的状态
	state := a
	fmt.Println(state)
	state = b
	fmt.Println(state)
}

func main0402() {
	// iota 在第一行写上 iota 后面的内容则自动加一
	const (
		a = iota
		b
		c
		d
		e
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}

// 如果定义的枚举的常量卸载同一行，那么同一行的几个常量值相同，换一行值加一
func main0403() {
	const (
		a    = 10
		b, c = iota, iota
		d, e
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
