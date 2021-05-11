package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 为什么需要 context
var wg sync.WaitGroup

func f2(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("context")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP // 指定跳到指定位置
		default:
		}
	}
}

func f(ctx context.Context) {
	defer wg.Done()
	go f2(ctx) // 将 ctx 传递给下一级
	LOOP:
		for {
			fmt.Println("hello")
			time.Sleep(time.Millisecond * 500)
			select {
			case <-ctx.Done(): // 等待上级通知
				break LOOP // 指定跳到指定位置
			default:
			}
		}
}

func main() {
	// 使用 context 通知 子 goroutine 结束
	// 创造一个 context 根节点
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx) // 将 ctx 传递给下一级
	time.Sleep(5 * time.Second)
	cancel() // 通知子 goroutine 结束
	wg.Wait()
}
