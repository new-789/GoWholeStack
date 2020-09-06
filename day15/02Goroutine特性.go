package main

import (
	"fmt"
	"time"
)

func main0201() {
	go func() { // 创建一个子 go 程
		for i := 0; i < 5; i++ {
			fmt.Println("------ I'm goroutine ------")
			time.Sleep(time.Second * 1)
		}
	}()

	for i := 0; i < 5; i++ { // 主 go 程
		fmt.Println("------- I'm main ------")
		if i == 2 {
			break
		}
	}
}
