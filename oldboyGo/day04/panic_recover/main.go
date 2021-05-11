package main

import "fmt"

// panic 和 recover

func f1() {
	defer func() {
		err := recover() // 收集当前程序中的错误情况
		fmt.Println("panic 放开那女孩，让我来", err)
		fmt.Println(err)
	}()
	panic("出了严重的故障！！！")
	fmt.Println("f1")
}

func f2() {
	fmt.Println("f2")
}

func main() {
	f1()
	f2()
}
