package main

import (
	"fmt"
	"strings"
)

// Contains 模糊查找使用示例
func main0101() {
	str := "hello world"
	// Contains 在一个字符串中查找另一个字符串是否出现，常用于模糊查找,返回值为 bool 类型
	value := strings.Contains(str, "llo")
	if value { // 将返回的 bool 值作为条件判断
		fmt.Println("找到了")
	} else {
		fmt.Println("未找到")
	}
}

// Join 字符串拼接使用示例
func main0102() {
	// Join 将一个字符串类型的切片拼接成一个字符串
	var s []string = []string{"1234", "5678", "91011", "1213", "1415"}
	str := strings.Join(s, "_")
	fmt.Println(str)
}

// Index 查找字符串下标使用示例
func main0103() {
	str := "踏破铁鞋无觅处，得来全不费工夫"
	// 在一个字符串中查找另一个字符是否出现，返回值为整型，如果找不到返回值为 -1
	index := strings.Index(str, "不费")
	fmt.Println(index)
}

// Repeat 重复一个字符串使用示例
func main0104() {
	// 将一个字符串重复 n 次，n 取值范围大于等于 0
	value := strings.Repeat("go", 3)
	fmt.Println(value)
}

// Replace 替换字符使用示例
func main0105() {
	str := "我擦你大爷"
	// 字符串替换元素，参数：字符串，被替换内容，替换内容， 替换次数 小于 0 则全部替换
	value := strings.Replace(str, "擦", "*", 1)
	value = strings.Replace(value, "大爷", "**", -1)
	fmt.Println(value)
}

// Split 字符串切割示例
func main0106() {
	str := "www.baidu.com"
	// 字符串切割，返回值是一个 []string
	value := strings.Split(str, ".")
	fmt.Println(value)
}

// Trim 去除字符串首尾内容示例
func main0107() {
	str := "=========are you ok?========="
	// 去掉字符串首尾的 =
	value := strings.Trim(str, "=")
	fmt.Println(value)
}

// Fields 去除字符串中的空格示例
func main0108() {
	str := "     Are         you      ok       ?     "
	// 去除字符串中的空格,返回返回有效字符串的类型切片
	s := strings.Fields(str)
	fmt.Println(s)
	// 将字符串类型的切片中的内容使用空格进行拼接
	result := strings.Join(s, " ")
	fmt.Println(result)
}
