package main

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

// AddBlock 5. 添加区块
func (b *BlockChain)AddBlock(data string) {
	// 获取最后一个区块作为当前区块的前区块哈希
	lastBlock := b.blocks[len(b.blocks) - 1]
	prevHash := lastBlock.Hash
	// a. 创建新的区块
	block := NewBlock(data, prevHash)
	// b. 添加到区块链数组中
	b.blocks = append(b.blocks, block)
}
