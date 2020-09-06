package main

import (
	"fmt"
	"time"
)

func sing() {
	for i := 0; i < 5; i++ {
		fmt.Println("赵四正在鬼哭狼嚎 《捉泥鳅》.......")
		time.Sleep(time.Microsecond * 100)
	}
}

func dance() {
	for i := 0; i < 5; i++ {
		fmt.Println("赵四又来跳 《赵四街舞》 舞蹈了.........")
		time.Sleep(time.Microsecond * 100)
	}
}

func main0101() {
	go sing()
	go dance()
	for {

	}
}
