package main

import (
	"fmt"
	"strings"
)

// 闭包案列
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		// 如果 name 参数不以 suffix 结尾，则加上后缀
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) // test,jpg
	fmt.Println(jpgFunc("hehe.jpg"))
	fmt.Println(txtFunc("test")) // test.html
	fmt.Println(txtFunc("hehe.txt"))
}
