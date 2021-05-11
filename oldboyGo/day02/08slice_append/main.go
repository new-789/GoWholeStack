package main

import "fmt"

// 切片扩容 append()，且切片追加元素
func main() {
	s1 := []string{"北京", "上海", "南宁"}
	fmt.Printf("s1:%v s1-len(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1))
	// s1[3] = "广州" // 错误的写法，会导致编译错误，索引越界
	// fmt.Println(s1)

	// 调用 append 函数必须用原切片变量接收返回值
	s1 = append(s1, "成都") // append 追加元素，原底层数组放不下的时候，Go 底层会在底层新建一个数组用来保存原数组中的值和新增的值，原数组则丢弃
	fmt.Println(s1)
	fmt.Printf("s1:%v s1-len(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "杭州", "广州")
	fmt.Printf("s1:%v s1-len(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1))
	ss := []string{"武汉", "长沙", "苏州"}
	s1 = append(s1, ss...) // ... 表示拆开切片
	fmt.Printf("s1:%v s1-len(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1))
}
