package main

import (
	"fmt"
	"strings"
)

func contFunc(str string) (m map[string]int) {
	// strings.Fields，将字符串进行切分，默认按照空格进行拆分，返回一个字符串切片
	s := strings.Fields(str)
	// 创建一个用于存储 val 出现次数的 map
	m = make(map[string]int)
	for i := 0; i < len(s); i++ {
		// 判断 map 中的 key 是否存在，不存在则添加并设置 value 为 1，存在则设置 value 为 2， !has 表示取反
		if _, has := m[s[i]]; !has {
			m[s[i]] = 1
		} else {
			m[s[i]] = m[s[i]] + 1
		}
	}
	return
}

func main0801() {
	var str string = "I Love my work and I Love my family too"
	m := contFunc(str)
	for k, v := range m {
		fmt.Printf("%s: %d\n", k, v)
	}
}
