package main

import "fmt"

// map 作为函数参数、返回值，map 作为函数参数默认属于传引用
func mapDel(m map[int]string, k int) {
	delete(m, k) // 删除 m 中键值为 k 的元素
}

func main0701() {
	m := map[int]string{
		9:   "我曾经",
		900: "等过你",
		200: "因为我也相信",
	}
	fmt.Println("before delete map m = ", m)
	mapDel(m, 9)
	fmt.Println("after delete map m = ", m)
}
