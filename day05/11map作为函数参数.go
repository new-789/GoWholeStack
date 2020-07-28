package main

import "fmt"

func demo(m map[int]string) {
	m[102] = "郑源"
	fmt.Printf("添加内容后的长度：%d\n", len(m))
	delete(m, 101) // 删除 map 参数中的值
	fmt.Printf("删除内容后的长度：%d\n", len(m))
}
func main1101() {
	m := map[int]string{101: "武大"}
	// map 作为参数传递时属于地址传递(引用传递)，意味着我们可以在接收map参数的函数中修改map中元素的值
	demo(m)
	fmt.Println(m)
}

func main1102() {
	m := map[int]string{101: "唐僧"}
	fmt.Printf("原长度：%d\n", len(m))
	demo(m)
}
