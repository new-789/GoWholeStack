package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Read 读取文件内容
func main0701() {
	// os.Open("读取文件的路径")  该函数仅以只读的方式打开文件
	file, err := os.Open("D:/test.txt")
	if err != nil {
		/*
			打开文件失败常见的集中原因
				1. 文件不存在
				2. 对需要打开的文件权限不足
				3. 文件打开的数量达到文件打开的上限65535个
		*/
		fmt.Println("Open file fail.....")
		return
	}
	// 关闭文件
	defer file.Close()

	b := make([]byte, 1024) // 定义读取文件内容的大小,及在内存中存放的位置
	file.Read(b)
	fmt.Println(b)
	fmt.Println(string(b))
	//for i := 0;i < len(b) ;i++  {
	//	if b[i] != '0'{
	//		fmt.Printf("%c", b[i])
	//	}
	//}
}

// ReadBytes 按行读取文件内容
func main0702() {
	fileObj, err := os.Open("D:/test.txt")
	if err != nil {
		fmt.Println("Open file err: ", err)
		return
	}

	defer fileObj.Close()

	// 创建切片缓冲区
	buf := bufio.NewReader(fileObj)
	// 读取一行内容,'\n' 表示每一行读取到换行结束，并将换行和文件内容一并存储到字节切片中
	b, er := buf.ReadBytes('\n')
	// 打印切片中字符的 ASCII 值
	//fmt.Println(b)

	if er != nil {
		fmt.Println("读取文件错误")
		return
	}

	// 将切片转成 string 打印汉字
	fmt.Println(string(b))
	b, _ = buf.ReadBytes('\n')
	fmt.Println(string(b))
}

// Read() 读取整个文件内容
func main0703() {
	Fobt, err := os.Open("D:/a.txt")
	if err != nil {
		fmt.Println("文件打开错误:", err)
		return
	}

	defer Fobt.Close()

	b := make([]byte, 20)

	for {
		// 读取文件返回值为个数和错误信息
		n, err := Fobt.Read(b)
		// io.EOF 表示文件的结尾，值为 -1, EOF == end of file
		if err == io.EOF { // 如果错误等于 io.EOF 错误表示读到了文件结尾直接跳出循环
			break
		}
		fmt.Print(string(b[:n]))
	}
}

// ReadBytes() 读取整个文件内容
func main0704() {
	file, err := os.Open("D:/a.txt")
	if err != nil {
		fmt.Println("文件打开错误:", err)
		return
	}

	defer file.Close()
	// 创建缓冲区
	buf := bufio.NewReader(file)
	// 通过循环一行行的读取文件中的内容
	for {
		// 读取一行内容，返回 byte 类型的数据内容和错误信息
		b, er := buf.ReadBytes('\n')

		fmt.Print(string(b)) // 打印文件内容

		// 如果错误等于 io.EOF 错误表示读到了文件结尾直接跳出循环
		if er == io.EOF {
			break
		}
	}
}

// saskkasjh
