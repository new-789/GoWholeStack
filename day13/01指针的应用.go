package main

import "fmt"

//func test(m int)  {
//	var b int = 1000
//	b += m
//}
//
//func main() {
//	var a int = 10  // 定义一个普变量
//	// 定义一个指针变量，并将变量 a 指向的内存地址存储在指针变量 b 所对应的内存空间中
//	var b *int = &a
//
//	a = 100
//	fmt.Println("a = ", a)
//
//	test(10)
//
//	// 通过 *b 解引用或间接引用修改变量 a 的值，实际上是借助 a 变量的地址操作 a 对应的空间
//	*b = 250
//	fmt.Println("a = ", a)
//	fmt.Println("*b = ", *b)
//	a = 1000
//	fmt.Println("*b = ", *b)
//}

// 指针变量的内存存储
//func main()  {
//	var b *int  // 未被初始化的空指针
// 在 heap 上申请一片内存地址空间
//	b = new(int)
//	*b = 100
//	fmt.Printf("%d\n", *b)
//}

// 传引用和传值的区别
func swap(a, b int) {
	a, b = b, a
	fmt.Printf("swap a:%d, b:%d\n", a, b)
}

func swap2(x, y *int) {
	*x, *y = *y, *x
}

func main0101() {
	a, b := 10, 20
	// 传值(传变量值)
	swap(a, b)
	fmt.Printf("swap1：main a:%d, b:%d\n", a, b)
	// 传引用(传地址值)
	swap2(&a, &b)
	fmt.Printf("swap2：main a:%d, b:%d\n", a, b)
	/*
		无论是传值还是传引用，函数传参永远遵循一个规律：值传递，只不过传值和传引用传递的内容不同
			传值：实参会将自己的值(数据值)拷贝一份传给形参；
			传引用：实参会将自己的值(地址值)拷贝一份传给形参；
	*/
}
