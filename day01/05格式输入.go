package main

import "fmt"

func main0501() {
	//var a int
	// & 运算符 取地址运算符
	//fmt.Scan(&a)
	//fmt.Println(a)

	// 空格或回车作为接收结束
	//var str string
	//fmt.Scan(&str)
	//fmt.Println(str)

	// 接收两个数据
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	fmt.Println(s1)
	fmt.Println(s2)
}

func main0502() {
	var r float64
	var PI float64

	// 通过键盘获取半径
	fmt.Scan(&r, &PI)
	fmt.Printf("面积:%.2f\n ", PI*r*r)
	fmt.Printf("周长: %.2f\n", 2*PI*r)
}

func main0503() {
	var a int
	var b string
	fmt.Scanf("%3d%s", &a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
