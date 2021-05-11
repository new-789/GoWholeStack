package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

// tailf 的用法

func main() {
	fileName := "./my.log"
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的那个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,                                 //
	}

	// 通过上面的配置文件打开文件，打开文件后会自动一行行的读取文件内容，并将读到的文件内容存入 channel 中
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}
	var (
	// line *tail.Line
	// ok   bool
	)
	for {
		line, ok := <-tails.Lines // 从 channel 中一行行的读取数据
		if !ok {
			fmt.Printf("tail faile close reopen, filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Printf("line:%v\n", line.Text)
	}
}
