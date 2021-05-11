package main

import (
	"bufio"
	"fmt"
	"os"
)

// 获取用户输入时如果有空格
func useScan() {
	fmt.Print("请输入内容:")
	var s string
	fmt.Scanln(&s)
	fmt.Printf("你输入的内容的：%s\n", s)
}

// bufio 获取用户标准输入
func userScan() {
	var s string
	fmt.Print("请输入：")
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Println("你输入的内容是：", s)
}

func main() {
	// useScan()
	userScan()
}
