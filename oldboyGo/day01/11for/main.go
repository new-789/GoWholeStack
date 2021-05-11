package main

import "fmt"

// for 循环

func main() {
	// 基本格式
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// 变种 1
	// var i = 5
	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// 变种 2
	// var i = 5
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// 无限循环
	// for {
	// 	fmt.Println(123)
	// }

	// for range 循环语句
	// s := "hello沙河"
	// for i, v := range s {
	// 	fmt.Printf("%d--%c\n", i, v)
	// }

	// 打印 9*9 乘法表
	for i := 0; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d\t", j, i, j*i)
		}
		fmt.Println()
	}
}
