package main

import "fmt"

func main0401() {
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i, j)
		}
	}
}

func main0402() {
	// 外层控制行， 内层控制列
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
			if i == j {
				// 在嵌套循环中，break 会跳出本次循环，无法跳出外层循环
				break
			}
		}
		fmt.Println()
	}
}

// 作业打印出下面格式的内容

//     *           5    1
//    ***          4    3
//   *****         3    5
//  *******        2    7
// *********       1    9
//***********      0    11

func main0403() {
	// 控制行
	for i := 0; i <= 5; i++ {
		// 控制空格个数,根据行数减去
		for j := 0; j < 5-i; j++ {
			fmt.Print(" ")
		}
		// 控制 * 的数量
		for k := 0; k < i*2+1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
