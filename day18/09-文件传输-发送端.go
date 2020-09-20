package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func sendFile(conn net.Conn, filePath string) {
	// 以只读方式打开文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("os.Open error:", err)
		return
	}
	defer file.Close()

	// 从本地文件中循环读取文件内容并发送给接收端，读多少写多少
	buf := make([]byte, 4096)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完成。。。。。。。")
			} else {
				fmt.Println("file.Raed error:", err)
			}
			return
		}

		// 发送文件内容到接收端
		_, err1 := conn.Write(buf[:n])
		if err1 != nil {
			fmt.Println("conn.Write fileInfo error:", err1)
			return
		}
	}
}

func main() {
	// 获取命令行参数
	list := os.Args
	// 提取文件绝对路径
	if len(list) != 2 {
		fmt.Println("语法错误，语法格式应为：go run xxx.go 文件绝对路径")
		return
	}
	filePath := list[1]

	// 提取不包含路径的文件名
	fileInfo, err1 := os.Stat(filePath)
	if err1 != nil {
		fmt.Println("os.Stat error:", err1)
		return
	}
	fileName := fileInfo.Name()

	// 主动发起连接请求
	conn, err2 := net.Dial("tcp", "127.0.0.1:8004")
	if err2 != nil {
		fmt.Println("net.Dial error:", err2)
		return
	}
	defer conn.Close()

	// 发送文件名给接收端
	_, err3 := conn.Write([]byte(fileName))
	if err3 != nil {
		fmt.Println("conn.Write fileNme error:", err3)
		return
	}

	// 服务器接收端回发的数据 ok
	buf := make([]byte, 4096)
	n, err4 := conn.Read(buf)
	if err4 != nil {
		fmt.Println("conn.Read error:", err4)
		return
	}

	// 判断接收端的回执是否为 ok 是则借助 conn 开始写文件内容到接收端
	if string(buf[:n]) == "ok" {
		sendFile(conn, filePath)
	}
}
