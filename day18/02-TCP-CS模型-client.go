package main

import (
	"fmt"
	"net"
)

func main0201() {
	// 指定协议和服务器 IP + PORT 创建和服务器通信的套接字
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("net.Dial fialed:", err)
		return
	}
	defer conn.Close()

	// 主动写数据发送给服务端
	conn.Write([]byte("Are you Ready"))

	buf := make([]byte, 4096)
	// 接收服务器返回的数据
	n, er := conn.Read(buf)
	if er != nil {
		fmt.Println("conn.Read failed:", er)
		return
	}
	fmt.Printf("服务器回发：%s\n", string(buf[:n]))
}
