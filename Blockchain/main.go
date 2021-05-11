package main

import "fmt"

//Block 1. 定义区块链结构
type Block struct {
	//	 前区块哈希
	PrevHash []byte
	//	 当前区块哈希
	Hash []byte
	//	 数据
	Data []byte
}
//NewBlock 2. 创建区块
func NewBlock(data string, precHash []byte) *Block {
	return &Block{
		PrevHash: precHash,
		Hash:     []byte{}, // 先填空，后面计算,TODO
		Data:     []byte(data),
	}
}
//3. 生成哈希
//4. 引入区块链
//5. 添加区块
//6. 重构代码

func main() {
	block := NewBlock("比特币实现简单版本", []byte{})
	fmt.Printf("前区块哈希值：%x\n", block.PrevHash)
	fmt.Printf("当前区块哈希值：%x\n", block.Hash)
	fmt.Printf("区块数据：%s\n", block.Data)
}
