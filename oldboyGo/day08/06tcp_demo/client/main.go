package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp client

func main() {
	//1. 与 producer 端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:8086")
	if err != nil {
		fmt.Println(" client dial failed, err:", err)
		return
	}
	defer conn.Close()
	// 2. 发送数据
	inputRead := bufio.NewReader(os.Stdin)
	var tmp [128]byte
	for {
		fmt.Print("请说话：")
		inputInfo, _ := inputRead.ReadString('\n') // 读取用户输入
		if strings.ToUpper(inputInfo) == "Q" { // 为什么无法退出？
			return
		}
		_, err = conn.Write([]byte(inputInfo))
		if err != nil {
			fmt.Println("send data to producer failed, err:", err)
			return
		}
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("client recv failed, err:",err)
			return
		}
		fmt.Println("收到服务端发来的数据：", string(tmp[:n]))
	}
}
