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
	// 1. 拼装数据
	blockInfo := append(b.PrevHash, b.Data...)
	// 2. sha256 生成哈希值
	hash := sha256.Sum256(blockInfo)
	b.Hash = hash[:] // 给当前区块添加哈希值
}

// BlockChain 4. 引入区块链
type BlockChain struct {
	// 定义一个区块链数组
	blocks []*Block
}
// NewBlockChain 5. 定义一个区块链
func NewBlockChain() *BlockChain {
	// 创建一个创世块，并作为第一个区块添加到区块链中
	genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesisBlock},
	}
}

// GenesisBlock 定义一个创世快
func GenesisBlock() *Block {
	return NewBlock("Go 5 期创世快", []byte{})
}

//5. 添加区块
//6. 重构代码

func main() {
	bc := NewBlockChain()
	//block := NewBlock("比特币实现简单版本", []byte{})
	for i, v := range bc.blocks {
		fmt.Printf("=========> 当前区块高度：%d ================\n", i)
		fmt.Printf("前区块哈希值：%x\n", v.PrevHash)
		fmt.Printf("当前区块哈希值：%x\n", v.Hash)
		fmt.Printf("区块数据：%s\n", v.Data)
	}
}