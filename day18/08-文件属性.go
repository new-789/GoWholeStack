package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取命令行参数,返回一个 list 数据
	list := os.Args

	if len(list) != 2 {
		fmt.Println("输入错误，格式应为：go run xxx.go 文件名")
		return
	}

	// 提取文件名
	path := list[1]

	// 获取文件属性
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat error:", err)
		return
	}

	fmt.Printf("文件名：%s\n", fileInfo.Name())
	fmt.Printf("文件大小: %d\n", fileInfo.Size())
}
