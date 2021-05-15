package main

import (
	"fmt"
	"os"
)

// 命令行测试Demo

func main() {
	len1 := len(os.Args)
	fmt.Println("cmd len:", len1)
	for i, cmd := range os.Args {
		fmt.Printf("arg[%d]: %s\n", i, cmd)
	}
}
