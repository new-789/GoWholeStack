package main

import "fmt"

func main0401() {
	// 定义一个数组
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("打印数组首地址：%p\n", &arr)
	fmt.Printf("打印数组中下标为 0 元素的地址：%p\n", &arr[0])
	// 定义指针指向数组,注：此处数组指针的类型必须和已有数组定义的元素个数一致
	var p *[5]int
	p = &arr // 将指针变量和数组建立关系
	fmt.Printf("指针存储的内存地址：%p\n", p)
	fmt.Println(p)
}

// 通过指针间接操作数组
func main0402() {
	arr := [5]int{1, 2, 3, 4, 5}
	var p *[5]int
	p = &arr
	// 通过指针间接修改数组下标为0的元素值为 123
	p[0] = 123
	// 通过指针变量打印输出数组的内容
	fmt.Println(*p)
	// 通过指针间接打印输出 arr 数组中下标为 1 的元素
	fmt.Println(p[1])
}

func main0403() {
	arr := [5]int{1, 2, 3, 4, 5}
	// 将指针与数组中下标元素为0 所指向的内存地址建立管理
	p := &arr[0]
	fmt.Printf("%T\n", p)
	// 将指针与数组所指向的内存地址建立关系
	p1 := &arr
	fmt.Printf("%T\n", p1)
	// 通过 for 循环访问数组中每一个元素
	//for i := 0;i < len(p) ;i++  {
	//	fmt.Println(p[i])
	//}

	//fmt.Printf("%p\n", p)
	//fmt.Printf("%p\n", &arr)
	//fmt.Printf("%p\n", &arr[0])
}
