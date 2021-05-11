package main

import (
	"fmt"
	"sort"
)

// 切片的练习
func main() {
	// 练习一
	var s = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		s = append(s, fmt.Sprintf("%v", i))
	}
	fmt.Println(s) // [0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
	fmt.Println(cap(s))

	var a = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a[:]) // 对切片进行排序
	fmt.Println(a)
}
