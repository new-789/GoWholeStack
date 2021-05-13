package main

import (
	"github.com/boltdb/bolt"
	"log"
)

// 区块的哈希值（即区块链的链条） Demo

// BlockChain 4. 引入区块链
type BlockChain struct {
	// 定义一个区块链数组
	db *bolt.DB
	tail []byte  // 用来存储最后一个区块的哈希值
}

const (
	blockChainDb = "blockChain.db"
	blockBucket = "blockBucket"
)
// 最后一个区块的哈希，从 bolt 数据库读出来的
var lastHash []byte


// NewBlockChain 5. 定义一个区块链
func NewBlockChain() *BlockChain {
	db, err := bolt.Open(blockChainDb, 0600, nil)
	if err != nil {
		log.Panicf("open blot DB failed, err:%v\n", err)
	}
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		// 返回为空则说明没有获取到了 bucket
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panicf("create blockBucket failed, err:%v\n", err)
			}
			// 创建一个创世块，并作为第一个区块添加到区块链中
			genesisBlock := GenesisBlock()
			// 写数据, Hash 作为 key, block 作为字节流(区块的数据)
			bucket.Put(genesisBlock.Hash, genesisBlock.toByte())
			// 更新 LastHash 对应的最后一个区块的哈希到数据库，方便我们查找最后一个区块的哈希
			bucket.Put([]byte("LastHashKey"), genesisBlock.Hash)
			lastHash = genesisBlock.Hash
		} else {
			// 读数据，并更新指向最后一个区块 key 的哈希值
			lastHash = bucket.Get([]byte("LastHashKey"))
		}
		return nil
	})
	return &BlockChain{db, lastHash}
}

// GenesisBlock 定义一个创世快
func GenesisBlock() *Block {
	return NewBlock("Go 5 期创世快", []byte{})
}

// AddBlock 5. 添加区块
func (b *BlockChain)AddBlock(data string) {
	/*
	// 获取最后一个区块作为当前区块的前区块哈希
	lastBlock := b.blocks[len(b.blocks) - 1]
	prevHash := lastBlock.Hash
	// a. 创建新的区块
	block := NewBlock(data, prevHash)
	// b. 添加到区块链数组中
	b.blocks = append(b.blocks, block)
	 */
}
