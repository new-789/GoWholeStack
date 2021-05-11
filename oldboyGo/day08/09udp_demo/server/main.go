package main

import (
	"fmt"
	"net"
)

// UDP SERVER

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP: net.IPv4(127, 0,0,1),
		Port: 8089,
	})
	if err != nil {
		fmt.Println("udp listen failed, err:", err)
		return
	}
	defer conn.Close()
	// 不需要建立连接，直接收发数据
	var tmp [1024]byte
	for {
		n, addr, err := conn.ReadFromUDP(tmp[:])
		if err != nil {
			fmt.Println("read from udp failed, err:",err)
			return
		}
		fmt.Printf("收到：%v 客户端数据：%v\n", addr, string(tmp[:n]))
		msg := "发送给客户端的数据"
		// 发送数据时第二个参数必须指定对方地址
		_, err = conn.WriteToUDP([]byte(msg),addr)
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			return
		}
	}
}
