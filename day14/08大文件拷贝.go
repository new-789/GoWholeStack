package main

import (
	"fmt"
	"io"
	"os"
)

func main0801() {
	// 打开需要拷贝的文件
	srcFile, err := os.Open("E:/迅雷下载/cn_windows_10_business_editions_version_2004_updated_july_2020_x64_dvd_f4ed1845.iso")
	if err != nil {
		fmt.Println("open file failed:", err)
		return
	}
	defer srcFile.Close()

	// 创建写入的文件
	dstFile, err1 := os.Create("d:/windows_x64_2004.iso")
	if err1 != nil {
		fmt.Println("create file err:", err1)
		return
	}
	defer dstFile.Close()

	// 创建读文件的缓冲
	buf := make([]byte, 1024*4)
	// 循环读取数据并写入文件
	for {
		// 开始读取文件内容
		n, err := srcFile.Read(buf)
		if n == 0 { // n 等于 0 表示已经读到了文件结尾
			fmt.Println("读文件结尾了,已经读完了........")
		}
		if err != nil && err == io.EOF {
			fmt.Printf("File read end, n = %d\n", n)
			break
		}

		// 开始写入文件
		_, er := dstFile.Write(buf[:n]) // 读读少写多少
		if er != nil {
			fmt.Println("write file failed:", er)
		}

	}
}
