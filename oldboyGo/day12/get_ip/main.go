package main

import (
	"fmt"
	"net"
	"strings"
)

// GetOutboundIP get ip demo 获取本地对外 IP
func GetOutboundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String()) // x.x.x.x:53744
	ip = strings.Split(localAddr.IP.String(), ":")[0]
	return
}

func main() {
	ip, _ := GetOutboundIP()
	fmt.Println(ip)
}
