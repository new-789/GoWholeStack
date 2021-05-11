package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net/http producer

// f1 对应 url 的处理函数
func f1(rw http.ResponseWriter, r *http.Request) {
	// 一次读取整个文件内容，返回字节数据和错误信息
	b, err := ioutil.ReadFile("./test.html")
	if err != nil {
		rw.Write([]byte(fmt.Sprintf("%v", err)))
	}
	// 返回内容给浏览器页面
	rw.Write(b)
}

func f2(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL) // 获取客户端的 url
	fmt.Println(r.Method) // 获取客户端的请求方法
	fmt.Println(ioutil.ReadAll(r.Body)) // 获取客户端请求的数据
	/* 对于 get 请求，参数都放在 url 上（query param），请求体中没有数据,
	 如想要获取 get 请求中 url 中的参数使用 Query 方法，
	它能自动帮我们识别 get 请求 URL 中所带的参数,返回 map 格式数据，如下：*/
	queryParam := r.URL.Query()
	// 通过返回的 map 数据获取参数数据
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)

	w.Write([]byte("收到了!"))
}

func main() {
	// HandleFunc 处理请求地址中的 url 和处理函数
	http.HandleFunc("/posts/Go/15_socket", f1)
	http.HandleFunc("/test/", f2)
	// ListenAndServe 启动监听 http 服务
	err := http.ListenAndServe("127.0.0.1:8081", nil)
	if err != nil {
		fmt.Println("start producer failed, err:", err)
		return
	}
}
