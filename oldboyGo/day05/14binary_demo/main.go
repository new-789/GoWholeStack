package main

import "fmt"

// 二进制实际用途

const (
	red int = 4
	gren int = 2
	blue int = 1
)

// 111
// 最左边的 1 表示红色 100
// 中间的1表示绿色 010
// 最右边的1表示蓝色 001

func color(arg int) {
	fmt.Printf("%b\n", arg)
}

func main() {
	color(red|blue) // 101
	color(red|gren|blue) // 111
}