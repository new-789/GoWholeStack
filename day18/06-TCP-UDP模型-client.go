package main

import (
	"fmt"
	"net"
	"os"
)

func main0601() {
	// 创建 UDP 通信 Socket,注意此处参数一和 TCP 协议版本代码的不同
	udpConn, err := net.Dial("udp", "127.0.0.1:8003")
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return
	}
	defer udpConn.Close()

	fmt.Println("客户端启动完成，可以开始发送数据了................")
	// 使用 go 程获取键盘输入的数据，然后发送到服务端
	go func() {
		for {
			str := make([]byte, 4096)
			if n, err := os.Stdin.Read(str); err != nil {
				fmt.Println("os.Stdin.Read error:", err)
				continue
			} else {
				// 发送数据到服务器
				udpConn.Write(str[:n])
			}
		}
	}()

	// 读取客户端的数据
	buf := make([]byte, 4096)
	for {
		n, err1 := udpConn.Read(buf)
		if err1 != nil {
			fmt.Println("udpConn.Read error:", err1)
			return
		}
		fmt.Printf("客户端读到服务器返回的数据：%s\n", string(buf[:n]))
	}
}
