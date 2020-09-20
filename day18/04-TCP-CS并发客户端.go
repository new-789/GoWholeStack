package main

import (
	"fmt"
	"net"
	"os"
)

func main0401() {
	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return
	}
	defer conn.Close()

	// 获取用户键盘输入(stdin)， 将数据发送给服务器
	go func() {
		str := make([]byte, 4096)
		for {
			// 获取键盘输入的数据
			if n, err := os.Stdin.Read(str); err != nil {
				fmt.Println("os.Stdin.Read error:", err)
				continue
			} else {
				// 发送数据给服务器，读到多少数据写多少
				conn.Write(str[:n])
			}
		}
	}()
	// 打印输出服务器返回的数据
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("检测到服务器以关闭，客户端即将退出")
			return
		}

		if err != nil {
			fmt.Println("conn.Read error: ", err)
			return
		} else {
			fmt.Println("客户端读到服务器数据：", string(buf[:n]))
		}
	}
}
