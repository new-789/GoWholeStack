package main

import (
	"fmt"
	"strconv"
	"sync"
)

// Go 内置的 map 不是并发安全的
var lock sync.Mutex
var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

// 不安全版的 map 多 goroutine 访问示例
//func main() {
//	wg := sync.WaitGroup{}
//	for i := 0; i < 200; i++ {
//		wg.Add(1)
//		go func(n int) {
//			key := strconv.Itoa(n)
//			lock.Lock()
//			set(key, n)
//			fmt.Printf("k=:%s, v=:%d\n", key, get(key))
//			lock.Unlock()
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//}

// sync.Map 安全版 Map 使用示例
var m1 = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 2000; i ++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(i)
			// 给 sync.Map 中设置值方法，必输使用sync.Map 提供的 Store 方法设置
			m1.Store(key, n)
			// 从 sync.Map 中获取值方法，必须使用 sync.Map 提供的 Load 方法
			val,_ := m1.Load(key)
			fmt.Printf("k:= %s, value:=%d\n", key,val)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
