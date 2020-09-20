package main

import (
	"fmt"
	"net"
	"runtime"
	"strings"
)

func HandlerConnet(conn net.Conn) {
	defer conn.Close()

	// 获取连接客户端的 IP 地址
	addr := conn.RemoteAddr()
	fmt.Println(addr, "client link success~~!")

	// 循环读取客户端发送的数据
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		// 客户端通过发送一个关键字来主动告诉服务端我要关闭了，服务器可以断开连接了,注意该方法 nc 工具模拟的客户端发送的数据会带有一个 \n 换行符
		if "exit\n" == string(buf[:n]) || "exit\r\n" == string(buf[:n]) {
			fmt.Printf("服务器接收到客户端 %s 的退出请求，以断开连接", conn.RemoteAddr())
			return
		}

		// 通过判断 n 的值是否为 0 确定客户端是否关闭
		if n == 0 {
			fmt.Printf("服务器检测到客户端 %s 以关闭，断开连接", conn.RemoteAddr())
			return
		}

		if err != nil {
			fmt.Println("conn.Read error:", err)
			runtime.Goexit()
		}
		// 模拟服务器处理数据
		fmt.Println("read client data: ", string(buf[:n]))
		// 将客户端的数据转换为大写，并回写给客户端
		data := strings.ToUpper(string(buf[:])) // 将小写转换为大写
		conn.Write([]byte(data))                // 回写数据给客户端
	}
}

func main0301() {
	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Listen error:", err)
		return
	}
	defer listener.Close()

	// 通过循环不断监听等待客户端连接请求实现并发
	for {
		fmt.Println("server start success, wait clinet link..................")
		conn, err1 := listener.Accept()
		if err1 != nil {
			fmt.Println("listener.Accept erroe:", err1)
			return
		}

		// 具体完成服务端和客户端的数据通信
		go HandlerConnet(conn)
	}
}
