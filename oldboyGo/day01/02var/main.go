package main

import "fmt"

// 声明变量
// var name string
// var age int
// var isOk bool

// 批量声明变量
var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "理想"
	// name := "test"
	age = 16
	isOk = true
	// go 语言中变量声明必须使用，不适用会出现编译不过去的错误提示
	fmt.Printf("name:%s\n", name) // %s 占位符，使用 name 这个变量的值去替换占位符
	fmt.Println(age)              // 打印完指定的内容之后会在后面加一个换行符
	fmt.Print(isOk)               // 在终端输出需要打印的内容
	fmt.Println()
	// 声明变量同时赋值
	var s1 string = "珠江"
	s1 := "10" // 同一个作用域不能同时声明一个变量
	// 类型推导，根据值判断该变量是什么类型
	var s2 = 20
	fmt.Println(s2)
	// 简短变量声明，只能在函数内使用
	s3 := "哈哈哈"
	fmt.Println(s3)
	// 匿名变量是一个特殊的变量：_，通常配合函数使用
}
