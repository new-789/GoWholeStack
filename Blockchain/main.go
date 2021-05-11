package main

import (
	"crypto/sha256"
	"fmt"
)

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

	block := &Block{
		PrevHash: precHash,
		Hash:     []byte{}, // 先填空，后面计算,TODO
		Data:     []byte(data),
	}
	block.SetHash()
	return block
}
// SetHash 3. 生成哈希
func (b *Block)SetHash() {
	// TODO
	// 1. 拼装数据，使用
	blockInfo := append(b.PrevHash, b.Data...)
	// 2. sha256
	hash := sha256.Sum256(blockInfo)
	b.Hash = hash[:]
}
//4. 引入区块链
//5. 添加区块
//6. 重构代码

func main() {
	block := NewBlock("比特币实现简单版本", []byte{})
	fmt.Printf("前区块哈希值：%x\n", block.PrevHash)
	fmt.Printf("当前区块哈希值：%x\n", block.Hash)
	fmt.Printf("区块数据：%s\n", block.Data)
}