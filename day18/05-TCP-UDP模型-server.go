package main

import (
	"fmt"
	"net"
	"time"
)

func main0501() {
	// 组织 UDP 地址结构, 指定服务器的 IP 和 Port, 返回一个服务器的地址结构
	srvAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8002")
	if err != nil {
		fmt.Println("net.ResolveUDPAddr error:", err)
		return
	}
	fmt.Println("服务器地址结构创建完成！！！")

	// 创建 UDP 通信的 Socket
	udpConn, er := net.ListenUDP("udp", srvAddr)
	if er != nil {
		fmt.Println("net.ListenUDP error:", er)
		return
	}
	defer udpConn.Close()
	fmt.Println("udp服务器通信 socket 创建完成，等待接收客户端数据.....................")

	// 读取客户端发送的数据
	buf := make([]byte, 4096)
	// 返回三个值：读取到的字节数，客户端的地址结构，error
	n, cltAddr, err1 := udpConn.ReadFromUDP(buf)
	if err1 != nil {
		fmt.Println("udpConn.ReadFromUDP error:", err1)
		return
	}

	// 模拟处理数据
	fmt.Printf("服务器读到 %s 发送过来的数据：%s\n", cltAddr, string(buf[:n]))

	// 提取系统当前并将时间转换为字符串类型
	dayTime := time.Now().String()
	// 回写数据到客户端
	_, err2 := udpConn.WriteToUDP([]byte(dayTime), cltAddr)
	if err2 != nil {
		fmt.Println("udpConn.WriteToUDP error:", err2)
		return
	}
}
