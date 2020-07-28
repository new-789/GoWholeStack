package main

import "fmt"

func main0901() {
	m := make(map[int]string, 1)
	m[0] = "悟空"
	// 在 map 中值允许重复
	m[1] = "八戒"
	m[2] = "沙僧"
	m[3] = "八戒"

	//fmt.Println(m[2])
	// 在 map 打印中如果出现没有定义的 map ，默认值为空
	//fmt.Println("---",m[5],"---")

	//v, ok := m[5]
	//if ok{
	//	fmt.Println(v)
	//}else {
	//	fmt.Println("key 不存在")
	//}

	for k, v := range m {
		// v, ok := m[k]
		fmt.Println(k, v)
	}

}
