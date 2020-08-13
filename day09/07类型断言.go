package main

import "fmt"

func main0701() {
	arr := make([]interface{}, 3)
	arr[0] = 123
	arr[1] = 3.14156
	arr[2] = "hello"
	arr = append(arr, 1.234, [3]int{1, 2, 3}, []int{1, 2, 3})

	for _, v := range arr {
		// 对数据 v 进行类型断言
		//value, ok := v.(int)
		//if ok {
		//	fmt.Println(value)
		//}

		// 类型断言，并结合 if 判断使用语法
		if value, ok := v.(int); ok {
			fmt.Println("整型数据：", value)
		} else if value, ok := v.(float64); ok {
			fmt.Println("浮点类型：", value)
		} else if value, ok := v.(string); ok {
			fmt.Println("字符串类型：", value)
		} else if value, ok := v.([3]int); ok {
			fmt.Println("数组类型：", value)
		} else if value, ok := v.([]int); ok {
			fmt.Println("切片类型：", value)
		}
	}
}
