package main

import "fmt"

func main0801() {
	// 定义 map 结构数据 map[keyType]valueType
	m := make(map[int]string, 1)
	// map 数据的赋值操作，赋值语法 map[key],key 是一个基本数据类型
	m[10] = "张三"
	m[2] = "李四"
	m[5] = "王二"

	// map 的遍历操作
	for k, v := range m {
		fmt.Printf("key：%d, value：%s\n", k, v)
	}
	//fmt.Println(m)
}

func main0802() {
	// 定义 map 数据结构语法二
	//m := map[string]int{
	//	"赵四":50,
	//	"赵四": 30,
	//}
	m := make(map[string]int, 1)
	m["赵四"] = 50
	_, ok := m["赵四"]
	// 通过验证 key 对应的 value 是否有值，根据条件作出相应的操作
	if ok { // 如果存在则打印该 key 已存在
		fmt.Println("该 key 已存在")
	} else { // 不存在则进行赋值操作
		m["赵四"] = 30
	}
	//m[20] = "王大拿"
	fmt.Println(m)
}
