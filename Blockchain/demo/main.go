package main

import "fmt"

// 计算区块链总量

func main() {
	// 1. 每 21 万个块减半
	total := 0.0  // 用来记录比特币总量
	// 2. 最初奖励50个
	currentReword := 50.0
	blockInterval := 21.0 // 比特币总量,单位是万
	// 3. 用一个循环来判断，累加
	for currentReword > 0 {
		// 每一个区间内的总量
		amount1 := blockInterval * currentReword
		// 除法在程序中效率低，所以此处除以 2 将其更改等价的乘法为 *= 0.5
		currentReword *= 0.5
		total += amount1
	}
	fmt.Println("比特币总量：", total, "万")
}
