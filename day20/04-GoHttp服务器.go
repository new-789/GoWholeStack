package main

import (
	"fmt"
	"net/http"
	"strings"
)

func MyHandle(w http.ResponseWriter, r *http.Request) {
	// w 写给客户端的数据内容
	w.Write([]byte("this is a Web server"))
	// r 从客户端读到的数据内容,并拆分获取请求头中的各种信息
	fmt.Printf("Header:%s\n", r.Header) // 包含整个请求头内容，map 存储类型
	fmt.Printf("Host:%s\n", r.Host)
	fmt.Printf("Method:%s\n", r.Method)
	fmt.Printf("URL:%s, %T\n", r.URL, r.URL.String())
	fmt.Printf("RemoteAddr:%s\n", r.RemoteAddr)
	fmt.Printf("Body:%s\n", r.Body)
	fileNme := strings.Split(r.URL.String(), "/")
	fmt.Println("fileNAME:", fileNme[1])
}

func main0401() {
	// 注册回调函数，该函数在客户端访问服务器时自动被调用
	http.HandleFunc("/qwe", MyHandle)

	// 绑定服务器监听地址
	http.ListenAndServe("127.0.0.1:8000", nil)
}
