package main

import (
	"fmt"
	"net"
	"runtime"
	"time"
)

func WriteToClient(udpConn *net.UDPConn, cltAddr *net.UDPAddr) {
	dataStr := time.Now().String()
	udpConn.WriteToUDP([]byte(dataStr), cltAddr)
}

func main0701() {
	// 组织一个 udp 地址结构，指定服务器的 IP+PORT
	srvAddr, err1 := net.ResolveUDPAddr("udp", "127.0.0.1:8003")
	if err1 != nil {
		fmt.Println("net.ResolveUDPAddr errpr:", err1)
		return
	}

	// 创建通信 Socket
	udpConn, err := net.ListenUDP("udp", srvAddr)
	if err != nil {
		fmt.Println("net.ListenUDP error:", err)
		return
	}
	defer udpConn.Close()
	fmt.Println("udp服务器通信 socket 创建完成，等待接收客户端数据.....................")

	// 读取客户端发送的数据
	readBuf := make([]byte, 4096)
	for {
		n, ctlAddr, err := udpConn.ReadFromUDP(readBuf)
		if err != nil {
			fmt.Println("udpConn.ReadFromUDP error:", err)
			runtime.Goexit()
		}
		// 模拟处理客户端发送的数据
		fmt.Printf("服务器读到 %s 客户端传送过来的数据：%s\n", ctlAddr, string(readBuf[:n]))

		// 开启 go 程回写数据给客户端
		go WriteToClient(udpConn, ctlAddr)
	}
}
