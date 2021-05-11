package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 按照指定顺序遍历 map
func main() {
	rand.Seed(time.Now().UnixNano())      // 初始化随机数种子
	scoreMap := make(map[string]int, 200) // 创建 map 类型遍历并初始化

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) // 生成 stu 开头的字符串
		value := rand.Intn(100)          // 生成 0~99 之间的随机数
		scoreMap[key] = value
	}

	// 取出 map 中的所有 key 存入切片
	keys := make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}

	// 对切片进行排序操作
	sort.Strings(keys)
	// 循环排序后切片中的 key 遍历 map
	for _, v := range keys {
		fmt.Println(v, scoreMap[v])
	}
}
