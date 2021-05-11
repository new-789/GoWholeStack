package main

import "fmt"

// // map 和 slice 组合
func main() {
	// 元素类型为 map 的切片
	var s1 = make([]map[int]string, 10, 10)
	// 没有对内部符 map 做初始化
	s1[0] = make(map[int]string, 1)
	s1[0][10] = "练习"
	fmt.Println(s1)

	// 值为切片类型的 map
	var m1 = make(map[string][]int, 10)
	m1["北京"] = []int{10, 20, 30}
	fmt.Println(m1)
}
