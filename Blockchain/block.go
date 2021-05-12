package main

import "crypto/sha256"

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
	// 1. 拼装数据
	blockInfo := append(b.PrevHash, b.Data...)
	// 2. sha256 生成哈希值
	hash := sha256.Sum256(blockInfo)
	b.Hash = hash[:] // 给当前区块添加哈希值
}