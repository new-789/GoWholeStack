package main

import "fmt"

// 关于 append 删除切片中的某个元素
func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := a[:]
	// 删掉索引为 1 的2
	s1 = append(s1[:1], s1[2:]...)
	fmt.Println(s1) // [1 3 4 5 6 7 8 9]
	fmt.Println(a)  // [1 3 4 5 6 7 8 9 9]
}
