package main

import (
	"fmt"
	"net"
	"os"
)

func recvFile(conn net.Conn, fileNme string) {
	// 按照文件名创建新文件
	file, err1 := os.Create(fileNme)
	if err1 != nil {
		fmt.Println("os.Create error:", err1)
		return
	}
	defer file.Close()

	// 从网络中读取数据写入本地文件
	buf := make([]byte, 4096)
	for {
		n, err2 := conn.Read(buf)
		// 如果接收到是数据长租为 0 则表示接收完毕或客户端已关闭
		if n == 0 {
			fmt.Println("文件内容接收完毕。。。。。。")
			return
		}
		if err2 != nil {
			fmt.Println("conn.Read fileInfo error:", err2)
			return
		}

		// 将文件内容写入到本地文件中
		file.Write(buf[:n])
	}
}

func main() {
	// 创建用于监听的 Socket
	listener, err1 := net.Listen("tcp", "127.0.0.1:8004")
	if err1 != nil {
		fmt.Println("net.Listen error:", err1)
		return
	}
	defer listener.Close()

	// 阻塞等待客户端连接
	fmt.Println("服务器创建通信 Socket 完成启动成功，等待客户端连接............")
	conn, err2 := listener.Accept()
	if err2 != nil {
		fmt.Println("listener.Accept error:", err2)
		return
	}
	defer conn.Close()

	// 读取客户端传送的文件名，并保存在指定位置
	readFileNmeBuf := make([]byte, 4096)
	n, err3 := conn.Read(readFileNmeBuf)
	if err3 != nil {
		fmt.Println("conn.Read fileName error:", err3)
		return
	}
	fileNme := string(readFileNmeBuf[:n])

	// 给发送端发送 ok 回执
	_, err4 := conn.Write([]byte("ok"))
	if err4 != nil {
		fmt.Println("conn.Write error:", err4)
		return
	}

	// 循环获取发送端发送的文件内容
	recvFile(conn, fileNme)
}
