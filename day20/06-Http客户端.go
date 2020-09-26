package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 获取服务器应答包内容
	res, err := http.Get("http://www.baidu.com/")
	if err != nil {
		fmt.Println("http.Get error:", err)
		return
	}
	// 此处一定要记得使用 Body 进行关闭
	defer res.Body.Close()

	// 查看应答包包含的内容
	fmt.Println("Header:", res.Header)
	fmt.Println("Request:", res.Request)
	fmt.Println("Proto:", res.Proto)
	fmt.Println("Status:", res.Status)
	fmt.Println("StatusCode:", res.StatusCode)
	fmt.Println("Close:", res.Close)
	fmt.Println("=================================================================")
	// 读取应答包中的  Body 内容
	buf := make([]byte, 4096)
	var result string
	for {
		n, err := res.Body.Read(buf)
		if n == 0 {
			fmt.Println("---------------Read Finish------------------")
			break
		}
		// 注意此处所使用的判断方法
		if err != nil && err != io.EOF {
			fmt.Println("res.Body.Read error:", err)
			return
		}
		result += string(buf[:n])
	}
	fmt.Printf("|%v|\n", result)
}
