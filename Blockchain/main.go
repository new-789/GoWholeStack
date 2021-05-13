package main

import "fmt"

func main() {
	bc := NewBlockChain()
	bc.AddBlock("111111111111111111111111111")
	bc.AddBlock("222222222222222222222222222")

	// 调用迭代器，返回每一个区块数据
	iterator := bc.NewIterator() // 创建一个迭代器
	for {
		// 返回区块并左移
		block := iterator.Next()
		fmt.Printf("====================================\n\n")
		fmt.Printf("前区块哈希值：%x\n", block.PrevHash)
		fmt.Printf("当前区块哈希值：%x\n", block.Hash)
		fmt.Printf("区块数据：%s\n", block.Data)
		// 判断如果单钱区块链的前一个区块切片长度为零则说明迭代完毕，则退出
		if len(block.PrevHash) == 0 {
			fmt.Printf("区块链遍历结束！")
			break
		}
	}
}
