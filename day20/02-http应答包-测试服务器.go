package main

import (
	"net/http"
)

// 回调函数
func Handler(w http.ResponseWriter, r *http.Request) {
	// w 表示写回给客户端(浏览器)是数据
	w.Write([]byte("helloWorld"))
	// r 表示从客户端(浏览器)读到的数据
}

func main0201() {
	// 注册回调函数，该回调函数会在服务器被访问时自动被调用
	http.HandleFunc("/itcast", Handler)
	// 绑定服务器监听地址
	http.ListenAndServe("127.0.0.1:8000", nil)
}
