package main

import "fmt"

func main0301() {
	// 切片定义, var 切片名 []数据类型
	var s []int
	//s[0]=1  定义的空切片无法使用下标的方式为切片添加元素
	fmt.Println("语法一", s)

	s1 := []int{1, 2, 3, 4}
	//s1[4] = 5  定义并初始化后的切片无法使用下标的方式为切片添加元素
	fmt.Println("语法二：", s1)

	s2 := make([]int, 5)
	s2[0] = 1
	s2[1] = 2
	s2[2] = 3
	s2[3] = 4
	s2[4] = 5
	//s2[5]=6  在超出创建切片指定的长度后依然无法使用下标的方式往切片中追加数据
	s2 = append(s2, 6, 7, 8, 9) // 通过 append 函数往切片中追加元素
	fmt.Println("语法三：", s2)
}

func main0302() {
	s := make([]int, 5)
	s[0] = 1
	s[1] = 2
	s[2] = 3
	s[3] = 4
	s[4] = 5
	// 遍历方式一
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	for i, v := range s {
		fmt.Println("下标为", i, "的元素为", v)
	}
}

func main0303() {
	var s []int = []int{1, 2, 3, 4, 5}

	s = append(s, 6, 7, 8)
	fmt.Println("长度", len(s))
	fmt.Println("容量", cap(s))

	s = append(s, 9, 10, 11)
	fmt.Println("长度", len(s))
	fmt.Println("容量", cap(s))
}

func main0304() {
	s := make([]int, 5)
	s[0] = 1
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

func main0305() {
	s := make([]int, 2)
	s = append(s, 1, 2, 3)

	for i := 0; i < cap(s); i++ {
		fmt.Println(s[i])
	}
}
