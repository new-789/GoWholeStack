package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// tcp 服务端实现
func process(conn net.Conn) {
	// 3. 与客户端通信
	reader := bufio.NewReader(os.Stdin)
	for {
		var tmp [128]byte
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("conn read failed, err:", err)
			return
		}
		fmt.Printf("读到客户端端发来的数据：%v\n", string(tmp[:n]))
		fmt.Print("请回话：")
		msg, _ := reader.ReadString('\n')
		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("producer conn write failed, err：", err)
			continue
		}
	}
}

func main() {
	// 1. 本地端口启动服务
	listen, err := net.Listen("tcp", "127.0.0.1:8086")
	if err != nil {
		fmt.Println("start tcp listen failed, err:", err)
		return
	}
	defer listen.Close()
	// 2. 等待别人跟我建立连接
	for {
		fmt.Println("等待客户端建立连接..............")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen Accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}