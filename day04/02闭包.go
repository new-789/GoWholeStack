package main

import "fmt"

func test1(a int) {
	a++
	fmt.Println(a)
}
func main0201() {
	a := 0
	for i := 0; i < 5; i++ {
		test1(a)
	}
}

// 1.定义一个函数，该函数的返回值为一个匿名函数
func test2() func() int {
	a := 0
	// 2.在 test2 主函数中返回一个匿名函数，该匿名函数返回一个 int 类型的数据
	return func() int {
		a++
		return a
	}
}

func main0202() {
	/* 3.将 test2 函数的返回值赋值给 f，由于 test2 返回值也是一个函数类型的返回值，所以此时 f 为一个函数类型的变量，
	  	 通过该变量就可以访问 test2 函数中的匿名函数，就实现了在一个函数中调用其它函数中的匿名函数；
		 注意：该语法与 f := test2 语法的区别，该语法则表示将 test2 函数类型赋值给变量 f
	*/
	f := test2()
	for i := 0; i < 5; i++ {
		// 4.使用 f 调用 test2 函数中的匿名函数，使用 v 来接收匿名函数的返回值
		v := f()
		fmt.Println(v)
	}
}
