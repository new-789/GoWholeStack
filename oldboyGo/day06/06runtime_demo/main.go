package main

import (
	"fmt"
	"path"
	"runtime"
)

// runtime.Caller

func getInfo(n int) {
	pc, file, line, ok := runtime.Caller(n)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(file)            // 06runtime_demo/mian.go
	fmt.Println(path.Base(file)) // main.go
	fmt.Println(line)            // 23
}

func f(n int) {
	getInfo(n)
}

func main() {
	f(2)
}
