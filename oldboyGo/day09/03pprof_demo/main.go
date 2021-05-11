package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

// 一段有问题的代码
func logicCode() {
	var c chan int  // nil
	for {
		select {
		case v := <- c: // 阻塞
			fmt.Printf("recv from chan ,value:%v\n", v)
		default:
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	var isCPUPprof bool // 是否开启 CPU profile 标志位
	var isMemPprof bool // 是否开启 内存 profile 标志位
	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()

	if isCPUPprof {
		f1, err := os.Create("./cpu.pprof") // 创建一个 cpu.pprof 文件
		if err != nil {
			fmt.Println("os create cpu pprof failed, err:", err)
			return
		}
		_ = pprof.StartCPUProfile(f1)
		defer func() {
			pprof.StopCPUProfile()
			_ = f1.Close()
		}()
	}
	for i := 0; i < 8; i++ {
		go logicCode()
	}
	time.Sleep(time.Second * 20)

	if isMemPprof {
		f2, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Println("create mem pprof failed, err", err)
			return
		}
		_ = pprof.WriteHeapProfile(f2)
		_ = f2.Close()
	}
}
