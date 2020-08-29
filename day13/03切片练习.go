package main

import "fmt"

// 切片练习一
func dupRem(old []string) (new []string) {
	new = make([]string, 0)
	for _, v := range old {
		if v != "" {
			new = append(new, v)
		}
		// 取到空字符串，不作为
	}
	return
}

// 不使用 append 方法实现去除切片中的空字符串,直接在原串上操作
func noEmpty(data []string) []string {
	i := 0
	for _, v := range data {
		if v != "" {
			data[i] = v
			i++
		}
	}
	return data[:i]
}

func main0301() {
	old := []string{"red", "", "black", "", "", "pink", "blue"}
	new := noEmpty(old)
	fmt.Println(new)
}
