package main

import (
	"fmt"
	"strings"
)

func main0401() {
	var str string = "I Love You and family You Love Mi"

	// Split 指定字符串分割使用示例
	/*
		ret := strings.Split(str, " ")
		fmt.Printf("ret= %v\nType= %T\n", ret, ret)
	*/

	// Fields 默认按空格分割使用示例
	/*
		ret := strings.Fields(str)
		fmt.Printf("ret= %v\nType= %T\n", ret, ret)
	*/

	// HasSuffix 判断字符串结束标记使用示例
	/*
		ret := strings.HasSuffix(str,"i")
		fmt.Printf("ret= %v\nType= %T\n", ret, ret)
	*/

	// HasPrefix 判断字符串起始标记使用示例
	ret := strings.HasPrefix(str, "I")
	fmt.Printf("ret= %v\nType= %T\n", ret, ret)
}
