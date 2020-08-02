package main

import "fmt"

func main0101() {
	// 定义 int 类型变量 a 并指定初始化值为10
	var a int = 100
	// 定义指针指针类型变量，指向 a 的地址，注指针类型变量的类型需要和变量 a 的类型一致
	var p *int
	// 将 a 变量的地址指向指针变量 p
	p = &a
	fmt.Println("指针变量存储的内容:", p)

	// 通过指针变量间接访问变量对应的内存空间
	fmt.Println("通过指针变量间接访问变量的内存空间", *p)

	// 通过指针变量间接访问变量对应的内存空间，并修改变量所对应的值
	*p = 123
	fmt.Println("修改了变量 a 之后的值：", a)
}

// 使用自动推导类型定义一个指针变量
func main0102() {
	a := 10
	p := &a
	*p = 123

	fmt.Println("修改了变量 a 之后的值：", a)
}

func main0103() {
	// 定义指针变量
	var p *int
	//*p = 129  // err
	//p = 0xc042058080  野指针
	*p = 128
	fmt.Println(*p)
}
