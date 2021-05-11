package main

import (
	"fmt"
	"net"
)

// udp client

func main() {
	conn, err := net.DialUDP("udp",nil, &net.UDPAddr{
		IP: net.IPv4(127,0,0,1),
		Port: 8089,
	})
	if err != nil {
		fmt.Println("udp DailUdp failed, err:", err)
		return
	}
	defer conn.Close()
	msg := "发送给服务端的数据"
	_, err = conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("udp write failed, err:", err)
		return
	}
	var data [1024]byte
	n,addr,err := conn.ReadFromUDP(data[:])
	if err != nil {
		fmt.Println("read from producer udp failed, err:",err)
		return
	}
	fmt.Printf("收到:%v 服务端数据：%v\n",addr, string(data[:n]))
}
