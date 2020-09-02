package main

import (
	"fmt"
	"os"
)

// create 创建文件示例
func main0501() {
	file, err := os.Create("e:/GoFileTest/test.txt")
	if err != nil {
		fmt.Println("Create err：", err)
		return
	}
	defer file.Close()
	fmt.Println("Create file success")

}

// Open 打开文件操作
func main0502() {
	file, err := os.Open("e:/GoFileTest/test.txt")
	if err != nil {
		fmt.Println("Open file err:", err)
		return
	}
	defer file.Close()
	if _, err := file.WriteString("对 Open 函数测试写操作"); err != nil {
		fmt.Println("write to file err:", err)
		return
	}
	fmt.Println("Open File success")
}

// OpenFile 打开文件操作
func main0503() {
	file, err := os.OpenFile("e:/GoFileTest/test.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("Open file err:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("openFile write file test"); err != nil {
		fmt.Println("Write to file err:", err)
		return
	} else {
		fmt.Println("Write to file success")
	}
	fmt.Println("open file success")
}
