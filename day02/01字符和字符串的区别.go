package main

import "fmt"

func main0101() {
	// 区别一：在内存中存储的值不同
	// 定义一个字符变量
	//var a byte = 'a'  // 存储的值为 'a'  --> 97
	// 定义一个字符串变量
	//var b string = "a"  // 保存的值为 "a" == 'a''\0'

	// 区别二：在字符串中 \0 表示结束符，格式化输出占位符 %s 遇到 \0 停止，
	// %s 遇到 \0 停止
	var c string = "hello world"
	fmt.Printf("%s", c)
}

func main0102() {
	var str1 string = "hello world"
	fmt.Println(len(str1))

	// 在 go 语言中一个汉字算作3个字符，为了和 Linux 系统统一处理
	var str string = "我和我的祖国"
	// 计算字符串的个数
	num := len(str)
	fmt.Println(num)
}
