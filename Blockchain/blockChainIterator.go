package main

import (
	"github.com/boltdb/bolt"
	"log"
)

// 迭代器

type BlockChainIterator struct {
	db *bolt.DB
	// 游标用于不断变换索引
	currentHashPointer []byte
}

//func NewIterator(bc *BlockChain) {
//
//}

func (b *BlockChain)NewIterator() *BlockChainIterator {
	return &BlockChainIterator{
		b.db,
		// 最初指向区块链的最后一个区块，随着 Next 的调用，不断变化
		b.tail,
	}
}

// Next 定义迭代器，
// Next 方法是属于迭代器的，但是迭代器是属于区块链的
// 1. 返回当前的区块
// 2. 指针前移
func (it *BlockChainIterator)Next() (block *Block) {
	it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("访问 bucket 错误，应该存在的请检查")
		}
		blockTmp := bucket.Get(it.currentHashPointer)
		// 解码动作
		block = DeSerialize(blockTmp)
		// 游标哈希左移，即将哈希指针指向当前区块链哈希值指针的前一个区块链哈希值指针
		it.currentHashPointer = block.PrevHash
		return nil
	})
	return
}