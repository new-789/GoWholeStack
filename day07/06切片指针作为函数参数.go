package main

import "fmt"

// 定义一个函数，并声明一个切片指针类型的参数
func test(s *[]int) {
	// 通过指针间接往切片中增加数据
	*s = append(*s, 4, 5, 6)
}

func main0601() {
	s := []int{1, 2, 3}
	/* 调用 test 函数，并通过取址符 & 获取切片 s 的内存地址作为实参进行传递
	切片指针传递，形参可以改变实参的值
	*/
	fmt.Printf("原切片中保存的地址值：%p\n", s)
	test(&s)
	// 打印修改后切片的内容
	fmt.Println(s)
	fmt.Printf("形参修改完切片内容后，切片中保存的地址值：%p\n", s)
}

///
