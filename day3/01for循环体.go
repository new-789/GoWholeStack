package main

import "fmt"

func main0101() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("1-100的和：%d\n", sum)
}

func main0102() {
	// 计算 1-100 之间所有偶数的和
	sum := 0
	//for i := 1; i <= 100; i++{
	//	if i % 2 == 0{
	//		sum += i
	//	}
	//}

	// 代码调优
	for i := 0; i <= 100; i += 2 {
		sum += i
	}
	fmt.Println(sum)
}

func main0103() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	//fmt.Println(i)
}

func main0104() {
	i := 0 // 模拟 for 循环语法中的表达式1
	for {
		if i > 5 { // 模拟 for 循环语法中的表达式2
			break // 跳出循环
		}
		fmt.Println(i)
		i++ // 模拟 for 循环语句中的表达式3
	}
}
