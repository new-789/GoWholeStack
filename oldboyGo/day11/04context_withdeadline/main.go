package main

import (
	"context"
	"fmt"
	"time"
)

// Context_WithDeadline Demo
func main() {
	d := time.Now().Add(5 * time.Second)
	// WithDeadline,用来指定超时时间，即超过第二个参数指定的时间后通知 子goroutine 结束
	ctx, cancel := context.WithDeadline(context.Background(), d)
	// 通知子 goroutine 该结束了,后台会调用 ctx.Done() 方法并往 chan 中加入一个空结构体
	/*
	尽管 ctx 会过期，但在任何情况下调用它的 cancel 函数都是很好的实践。
	如果不这样做，可能会使上下文及其父类存活的实间超过必要的时间。*/
	defer cancel()

	select {
	case <- time.After(time.Second * 1):
		fmt.Println("overslept")
	case <- ctx.Done():
		fmt.Println(ctx.Err())
	}
}
