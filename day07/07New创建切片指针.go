package main

import "fmt"

func main0701() {
	// 定义一个空切片指针类型
	var p *[]int
	fmt.Printf("空指针地址：%p\n", p)
	// 通过 new 创建切片指针空间，即使用该语法在堆区开辟一块空间用来存储切片的内容，类型为切片类型
	p = new([]int)
	fmt.Printf("创建空间后地址：%p\n", p)
	*p = append(*p, 1, 2, 3)

	//for i := 0;i < len(*p) ;i++  {
	//	fmt.Println((*p)[i])
	//}

	for i, v := range *p {
		fmt.Println("下标：", i, "具体的值：", v)
	}
}
