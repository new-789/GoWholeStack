package main

import "fmt"

type stu struct {
	name  string
	age   int
	score int
}

func main0501() {
	// 定义 map
	m := make(map[int]stu)
	// 对 map 数据结构体结构体类型的值进行赋值操作
	m[101] = stu{"方丈", 89, 100}
	m[102] = stu{"沙弥", 20, 60}

	fmt.Println(m)

	// 循环 map 打印每一个 value
	for k, v := range m {
		fmt.Println("key：", k, "value：", v)
	}
}

// map 值为 结构体切片应用实例
func main0502() {
	// 定义一个 map，并将结构体切片类型作为该 map 的值
	m := make(map[string][]stu)
	// 往 map 中添加数据，由于 value 是结构体切片类型的数据，所以使用 append 添加结构体中的数据
	m["王者荣耀班"] = append(m["战士"], stu{"孙尚香", 20, 80}, stu{"吕布", 38, 90}, stu{"恺", 25, 85})
	// 通过双层循环输出结构体中的数据
	for k, v := range m {
		for i, data := range v {
			fmt.Println("班级: ", k, "学员类别：", i, "学员信息：", data)
		}
	}
}
