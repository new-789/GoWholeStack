package main

import (
	"fmt"
	"net"
)

func main0101() {
	// 指定服务器通信协议、IP 地址和 Port，注意：listen 不是用来监听只是用来绑定 IP和port 创建监听的 Socket
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Printf("net.listen fialed:%s\n", err)
		return
	}
	defer listener.Close() // 执行完毕则关闭用于监听的 Socket

	fmt.Println("server wait for client request link............")
	// 阻塞客户端连接请求,成功建立连接返回用于通信的 Socket
	conn, er := listener.Accept()
	if er != nil {
		fmt.Printf("listener accept fialed:%s\n", er)
		return
	}
	defer conn.Close() // 执行完毕则关闭用于通信的 Socket

	fmt.Println("server and client success link!!!")
	buf := make([]byte, 4096)
	// 读取客户端发送过来的数据
	n, err1 := conn.Read(buf)
	if err1 != nil {
		fmt.Printf("conn.Read failed:%s\n", err1)
		return
	}

	// 处理数据  ---- 此处只用来打印
	fmt.Printf("server read client data:%s\n", string(buf[:n]))

	conn.Write(buf[:n]) // 读取到客户端多少数据，就原封不动的给返回去
}
