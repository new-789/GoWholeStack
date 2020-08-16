package main

import (
	"fmt"
	"io"
	"os"
)

func main0801() {
	// 打开文件,从一个文件拷贝到新的文件中
	file1, err := os.Open("E:/Go语言区块链全栈开发/基础课/01_Go语言基础第10天（异常处理和文件读写）/03视频/10读取文件.mp4")
	file2, er := os.Create("D:/10读取文件.mp4")
	if err != nil || er != nil {
		fmt.Println("拷贝文件失败")
		return
	}
	defer file1.Close()
	defer file2.Close()

	b := make([]byte, 1024*1024)
	for {
		// 读取文件内容,拷贝文件应使用块读取
		n, err := file1.Read(b)
		if err == io.EOF {
			break
		}
		// 保存文件内容
		_, er := file2.Write(b[:n])
		if er != nil {
			fmt.Println("文件拷贝错误")
			break
		}
	}
	fmt.Println("文件拷贝完成")
	// 上面 for 循环中的拷贝操作可以使用 io 包中的 Cope 函数一次性解决如下(但对于初学者来说不建议这么操作)
	//io.Copy(file1, file2)
}
