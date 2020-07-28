package main

import "fmt"

// 切片添加数据
func main0501() {
	var s []int                  // 定义一个空切片
	fmt.Println("空切片长度", len(s)) // 打印空切片长度
	fmt.Println("空切片容量", cap(s)) // 打印空切片容量
	fmt.Printf("空切片地址%p\n", s)   // 打印空切片地址

	s = append(s, 1, 2, 3, 4, 5)
	fmt.Println(s)
	fmt.Println("切片第一次追加数据后长度", len(s)) // 打印空切片长度
	fmt.Println("切片第一次追加数据后容量", cap(s)) // 打印空切片容量
	fmt.Printf("切片第一次追加内容后地址%p\n", s)   // 打印空切片地址

	s = append(s, 1, 2, 3, 4, 5)
	fmt.Println(s)
	fmt.Println("切片第二次追加数据后长度", len(s)) // 打印空切片长度
	fmt.Println("切片第二次追加数据后容量", cap(s)) // 打印空切片容量
	fmt.Printf("切片第二次追加内容后地址%p\n", s)   // 打印空切片地址
}

// 切片拷贝
func main0502() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 5)
	// 将 s1 切片中的数据拷贝到 s2 切片中，s2 中要有足够的容量,
	// 使用 copy 拷贝后 s1 和 s2 是两个独立的空间，修改其中一个切片中的内容不会对另一个切片产生影响
	copy(s2, s1)
	s2[2] = 100
	fmt.Println("s1 切片中的元素：", s1)
	fmt.Println("s2 切片中的元素：", s2)

	fmt.Printf("s1 切片地址: %p\n", s1)
	fmt.Printf("s2 切片地址: %p\n", s2)
}

func main0503() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 5)
	// 拷贝 s1 切片中下标从1开始到4结束之间的内容
	copy(s2, s1[1:4])
	fmt.Println(s2)
}
