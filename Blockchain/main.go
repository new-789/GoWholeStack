package main

import (
	"fmt"
)

func main() {
	bc := NewBlockChain()
	bc.AddBlock("三藏收了悟空为大徒弟，并封了 5000 的大红包")
	bc.AddBlock("三藏又收了悟能为二徒弟，并封了 2500 的大红包")
	for i, v := range bc.blocks {
		fmt.Printf("=========> 当前区块高度：%d ================\n", i)
		fmt.Printf("前区块哈希值：%x\n", v.PrevHash)
		fmt.Printf("当前区块哈希值：%x\n", v.Hash)
		fmt.Printf("区块数据：%s\n", v.Data)
	}
}