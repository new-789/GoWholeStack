package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
	"io/ioutil"
)
// 打开文件

func readFromFile1() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()
	// 读文件
	// var tmp = make([]byte, 128) // 指定读的长度
	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("文件读完了")
			return
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}

// 利用 bufio 包读文件一行行的读文件
func readFromFileByBufio() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed, err", err)
		return
	}
	defer file.Close()
	// 创建 read 读对象，NewReader 接收一个文件对象为参数
	read := bufio.NewReader(file)
	for {
		// ReadString 一行行的读取文件内容，接受的参数为字符类型的 \n
		line, err := read.ReadString('\n') // 此处是字符
		if err == io.EOF {  // 此错误为读到了文件末尾
			fmt.Println("文件读完了")
			return
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}
}

// ioutil 读文件操作，一次读取整个文件内容,ioutil 包存在于 io/ioutil 
func readFromFileByIoutil() {
	// ioutil 读文件操作方法，返回一个字节切片内容，和错误信息
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(ret))
}

func main() {
	// readFromFile1()
	// readFromFileByBufio()
	readFromFileByIoutil()
}