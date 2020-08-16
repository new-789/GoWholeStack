package main

import "fmt"

func test1() {
	fmt.Println("hello world1")
}

func test2() {
	//fmt.Println("hello world2")
	// 直接调用 panic 函数
	panic("hello world2")
}

func test3() {
	fmt.Println("hello world3")
}

func main0201() {
	test1()
	test2()
	test3()
}

// saskkasjh
