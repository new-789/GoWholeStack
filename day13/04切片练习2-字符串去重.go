package main

import "fmt"

func noSame(data []string) (newSlice []string) {
	newSlice = make([]string, 1)
	// 遍历原始切片字符串
	for _, v := range data {
		i := 0 // i 放在此处是为了保证相等跳出内循环之后每次都从新切片的第一个内容开始比较
		// 比较取出的 v 在 new 切片中是否存在
		for ; i < len(newSlice); i++ {
			if v == newSlice[i] {
				break
			}
		}
		// 该判断用于确保判断两个值是否相等的内部循环判执行完毕
		if i == len(newSlice) {
			newSlice = append(newSlice, v)
		}
	}
	return
}

func main0401() {
	s := []string{"red", "yellow", "black", "red", "pink", "yellow", "blue", "pink", "blue"}
	newSlice := noSame(s)
	fmt.Println(newSlice)
}
