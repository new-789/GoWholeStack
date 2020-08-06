package main

import "fmt"

//func add(a, b int) int  {
//	return a + b
//}

/*
 type 的两个作用
	1. 定义函数类型
	2. 为已存在的数据类型起别名
*/

type Int int // 为已存在的数据类型起别名

func (a Int) add(b Int) Int {
	return a + b
}

func main0601() {
	// 根据数据类型绑定方法
	var a Int = 10
	result := a.add(20)

	fmt.Println(result)
}
