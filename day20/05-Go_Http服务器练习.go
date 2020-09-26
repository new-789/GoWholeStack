package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func OpenSendFile(w http.ResponseWriter, fileName string) {
	// 拼接路径
	PathFileName := "E:/CodingFiles/GolangCode/test/" + fileName
	// 打开文件
	if file, err := os.OpenFile(PathFileName, os.O_RDONLY, 5); err != nil {
		w.Write([]byte("404 error!\r\nNo such file or directory"))
		return
	} else {
		defer file.Close()
		// 读取文件并返回给客户端
		buf := make([]byte, 4096)
		for {
			n, err := file.Read(buf)
			if err == nil && err == io.EOF {
				fmt.Println("file.Read error:", err)
				return
			}
			w.Write([]byte(buf[:n]))
		}
	}
}

func myHandle(w http.ResponseWriter, r *http.Request) {
	// 获取客户端 Url 并将其转换为 string 类型
	cltUrlName := r.URL.String()
	OpenSendFile(w, cltUrlName)
}

func main0501() {
	// 注册回调函数
	http.HandleFunc("/", myHandle)
	// 绑定服务器监听地址
	http.ListenAndServe("127.0.0.1:8000", nil)
}
