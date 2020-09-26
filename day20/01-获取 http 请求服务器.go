package main

import (
	"fmt"
	"net"
	"os"
)

// 出错处理函数封装
func errFunc(err error, info string) {
	if err != nil {
		fmt.Println(info, err)
		// return   // 返回当前函数调用
		// runtime.Goexit()  // 结束当前 go 程
		os.Exit(1) // 将当前进程结束，出错结束则返回一个非零值
	}
}

func main0101() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	errFunc(err, "net.Listen error:")
	defer listener.Close()

	fmt.Println("服务器启动完毕...................")
	conn, err1 := listener.Accept()
	errFunc(err1, "listener.Accept error:")
	defer conn.Close()

	buf := make([]byte, 4096)
	n, err2 := conn.Read(buf)
	if n == 0 {
		fmt.Println("客户端以关闭...............")
		return
	}
	errFunc(err2, "conn.Read error:")
	fmt.Printf("| %s|\n", string(buf[:n]))
}
