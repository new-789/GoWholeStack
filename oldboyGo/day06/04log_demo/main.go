package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// log_demo

func main() {
	file, err := os.OpenFile("./test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	// 设置日志输出的位置，os.Stdout 表示输出到终端，此处设置为输出到文件
	log.SetOutput(file)
	for {
		log.Println("这是一个测试日志")
		time.Sleep(3 * time.Second)
	}
}
