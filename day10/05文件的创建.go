package main

import (
	"fmt"
	"os"
)

func main0501() {
	file, err := os.Create("D:/test.txt")
	// 如果 err 内容不为空则表示文件创建失败
	if err != nil {
		fmt.Println("文件创建失败")
		return // 如果 return 出现在主函数中 表示程序的结束
	}

	// 延迟调用关闭文件,若不及时关闭文件则会造成以下两个问题
	/*
		1. 占用内存和缓冲区
		2. 文件打开上限用完，一个应用程序文件的打开上限为 65535 个
	*/
	defer file.Close()

	fmt.Println("文件创建成功")
}
