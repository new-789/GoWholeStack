package main

import (
	"fmt"
	"github.com/GoWholeStack/oldboyGo/day08/08nianbao_jiejue/proto"
	"net"
)

// 粘包 client
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err:", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := "Hello, Hello, How are you?"
		// 调用协议编码数据
		tmp, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("Proto encode msg failed, err:", err)
			return
		}
		conn.Write(tmp) // 发送数据
	}
}
