package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与 json
type person struct {
	// `json 表示反射格式为 json。 db：表示为数据库。ini：表示为 ini 配置文件`
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "悟能",
		Age:  999,
	}
	// 1. 序列化：将 go 语言中的结构体变量转变成 json 格式的字符串
	b, err := json.Marshal(p1) // 将 go 语言中的结构体转换成字符串
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("%v\n", string(b)) // 将序列化返回的字符切片强转为 string 类型

	// 2. 反序列化：将 json 格式的字符串转换成 go 语言中能够识别的结构体
	str := `{"name":"理想","age":"18"}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) // 传指针是为了能在 Unmarshal 函数内部修改 p2 的值
	fmt.Printf("%#v\n", p2)
}
