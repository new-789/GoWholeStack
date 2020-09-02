package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main0701() {
	file, err := os.OpenFile("C:/Users/zhufeng/Desktop/test/test.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	defer file.Close()

	// 按行读文件操作
	reader := bufio.NewReader(file)
	// 循环读取文件
	for {
		// 表示读到 \n 结束
		buf, er := reader.ReadBytes('\n')
		// io.EOF  为文件读完之后的一个标记,此处判断文件是否已经读取完毕
		if er != nil && er == io.EOF {
			fmt.Println("文件读取完毕")
			return
		} else if er != nil {
			fmt.Println("readBytes err:", er)
			return
		}
		fmt.Print(string(buf))
	}
}
