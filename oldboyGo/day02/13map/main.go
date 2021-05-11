package main

import "fmt"

// map
func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)        // 还没有初始化，没有在内存中开辟空间
	m1 = make(map[string]int, 10) // 初始化map,在对 map 初始化时尽量估算好 map 的容量，避免在程序运行期间动态扩容
	m1["理想"] = 20
	m1["张三"] = 28
	fmt.Println(m1)
	fmt.Println(m1["理想"])
	fmt.Println(m1["娜娜"]) // 如果 key 不存在，则返回对应 value 类型的零值

	// 判断 map 中是否存在某个值方法，约定俗成使用 OK 接收返回的布尔值
	v, ok := m1["娜娜"]
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(v)
	}

	// map 的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	// 只遍历 key
	for k := range m1 {
		fmt.Println(k)
	}
	// 只遍历 v
	for _, v := range m1 {
		fmt.Println(v)
	}
	// 删除 map 中元素
	delete(m1, "张三")
	fmt.Println(m1)
	delete(m1, "沙河") // 删除不存在的 key
}
