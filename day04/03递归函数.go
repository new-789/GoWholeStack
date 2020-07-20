package main

import "fmt"

func test3(a int) {
	if a == 0 {
		return
	}
	a--
	fmt.Println(a) // 在此处打印是在递归完成后，销毁数据出栈时依次打印销毁前的数据
	test3(a)
}
func main0301() {
	test3(5)
}

var sum int = 1

func test4(n int) {
	if n == 1 {
		return
	}

	test4(n - 1)
	sum *= n
}

func main0302() {
	n := 5
	test4(n)
	fmt.Println(sum)
}
