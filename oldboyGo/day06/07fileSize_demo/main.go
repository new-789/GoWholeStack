package main

import (
	"fmt"
	"os"
)

// 文件对象的类型
// 获取文件对象的详细信息

func main() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	// 文件对象的类型
	fmt.Printf("%T\n", fileObj)
	// 获取文件对象的详细信息
	fileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Println("file stat failed. err:", err)
		return
	}
	fmt.Printf("文件大小是：%dB\n", fileInfo.Size())
}
