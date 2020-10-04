package main

import (
	"fmt"
	rege "regexp"
)

func main0301() {
	str := "abc a7c mfc cat 8ca aMc szc cba"
	// 编译解析正则
	reg := rege.MustCompile(`a[^0-9a-z]c`) // `` 表示使用原生字符串
	// 从指定字符串中提取需要的信息
	SliceData := reg.FindAllStringSubmatch(str, -1)
	fmt.Println("SliceData:", SliceData)
}
