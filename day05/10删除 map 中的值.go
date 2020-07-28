package main

import "fmt"

func main1001() {
	m := map[int]string{101: "刘备", 102: "关羽", 103: "张飞", 104: "赵云", 105: "黄忠", 106: "武松"}

	// 打印数据类型
	fmt.Printf("%T\n", m)

	// 删除 map 数据 delete[map, key]
	delete(m, 106)
	// delete 删除不存在的值时，不会报错
	delete(m, 107)
	delete(m, 108)
	fmt.Println(m)
}
