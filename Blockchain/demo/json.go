package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	str1 := []string{"hello", "world", "!"}
	// 对一个字符串切片进行拼接
	res := strings.Join(str1, " ")
	fmt.Println(res)

	// 对一个二维切片通过一个一维切片进行拼接
	res1 := bytes.Join([][]byte{[]byte("hello"), []byte("world")}, []byte(""))
	fmt.Printf("res1:%s\n", res1)
}
