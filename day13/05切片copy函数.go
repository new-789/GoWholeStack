package main

import "fmt"

func main0501() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := data[8:] // [8,9]
	s2 := data[:5] //[0,1,2,3,4]

	copy(s2, s1) // 将 s1 切片中的内容按照对应位置 copy 到 s2 切片中
	fmt.Println("s2 = ", s2)
}

// 练习：删除下面 slice 中间的某个元素并保存原有的元素顺序，如:
func removeSlice(data []int, idx int) []int {
	copy(data[idx:], data[idx+1:])
	return data[:len(data)-1]
}

func main0502() {
	data := []int{5, 6, 7, 8, 9}
	afterData := removeSlice(data, 2)
	fmt.Println("afterData = ", afterData)
}
