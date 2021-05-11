package main

import (
	"bufio"
	"fmt"
	"github.com/GoWholeStack/oldboyGo/day08/08nianbao_jiejue/proto"
	"io"
	"net"
)

// 粘包server

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		// 调用自定义协议解码
		recvStr, err := proto.Decode(reader)
		if err == io.EOF {
			fmt.Println("数据读取完毕....")
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("收到 client 发来的数据：", recvStr)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
