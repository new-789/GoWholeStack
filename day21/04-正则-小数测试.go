package main

import (
	"fmt"
	re "regexp"
)

func main0401() {
	str := "3.14 123.123 .68 haha 1.0 abc 7. ab.3 66.6 123."

	// 解析编译正则表达式
	//reg := re.MustCompile(`[0-9]+\.[0-9]+`)
	reg := re.MustCompile(`\d?\.\d+`) // \d 同样可表示匹配整数 \D 表示匹配非整数字符

	// 从字符串中提取需要的信息
	sliceData := reg.FindAllStringSubmatch(str, -1)
	fmt.Println("sliceData:", sliceData)
}
