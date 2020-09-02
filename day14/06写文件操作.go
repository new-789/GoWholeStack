package main

import (
	"fmt"
	"io"
	"os"
)

func main0601() {
	f, err := os.OpenFile("e:/GoFileTest/test.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("open file err：", err)
		return
	}
	defer f.Close()
	fmt.Println("file open success, start write info to file............")
	// 文件写操作方式一, 按字符串写入
	/*
		if n, err := f.WriteString("helloWorld\r\n"); err != nil {
			fmt.Println("write file err:", err)
		} else {
			fmt.Println("write file success, Write info n = ", n)
		}
	*/

	// 文件写操作方式二，按位置写操作
	// 获取光标的偏移量
	if ret, err := f.Seek(-5, io.SeekEnd); err != nil {
		fmt.Println("obtain cursor offset err:", err)
	} else {
		// 打印输出获取到的光标偏移量
		fmt.Printf("current offset is:%d start write info to file......\n", ret)
		// 文件写操作方式三，按字节方式写操作
		msg := []byte{'2', '0', '2', '0'} // == []byte("2020")
		if n, er := f.WriteAt(msg, ret); er != nil {
			fmt.Println("Write msg to file err:", er)
			return
		} else {
			fmt.Println("write file success, write info num = ", n)
		}
	}
}
