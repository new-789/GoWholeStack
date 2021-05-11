package main

import "fmt"

// copy 切片的拷贝
func main() {
	s1 := []int{1, 3, 5}
	s2 := s1
	s3 := make([]int, 3)
	// 将一个切片 copy 到另一个切片时，前提是两个切片的长度必须一致
	copy(s3, s1)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	s1[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	// 删除切片中 的元素，将 s1 中索引为 1 的元素删除
	s1 = append(s1[:1], s1[2:]...)
	fmt.Println(s1)
	fmt.Println(cap(s1))
}
