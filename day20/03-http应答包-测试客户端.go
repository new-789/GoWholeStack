package main

import (
	"fmt"
	"net"
	"os"
)

func errFunc1(err error, info string) {
	if err != nil {
		fmt.Println(info, err)
		os.Exit(1)
	}
}

// 模仿浏览器
func main0301() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	errFunc1(err, "net.Dial error")
	defer conn.Close()
	// 组织 http 请求包，
	httpRequest := "GET /itcast HTTP/1.1\r\nHost:127.0.0.1:8000\r\n\r\n"
	_, Werr := conn.Write([]byte(httpRequest))
	errFunc1(Werr, "conn.Write error:")

	buf := make([]byte, 4096)

	n, Rerr := conn.Read(buf)
	if n == 0 {
		return
	}
	errFunc1(Rerr, "conn.Read error")

	fmt.Printf("|%s|", string(buf[:n]))
}
