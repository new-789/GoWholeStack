package main

import "fmt"

func main0201() {
	// 创建一个空指针变量
	var p *int
	// 在堆区创建指针空间，创建好的空间值为指定类型的默认值
	p = new(int)
	// 打印空指针内存空间中存储的内容
	fmt.Println("创建空指针空间后变量 P 对应的值: ", p)
	// 打印 p 指向空间的值
	fmt.Println("打印 p 指向空间的值: ", *p)
}
