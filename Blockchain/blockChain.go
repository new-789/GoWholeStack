package main

import (
	"bytes"
	"fmt"
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
func NewBlockChain(address string) *BlockChain {
	db, err := bolt.Open(blockChainDb, 0600, nil)
	if err != nil {
		log.Panicf("open blot DB failed, err:%v\n", err)
	}
	//defer db.Close() 此处不能关闭数据库，否则不能重复写入
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		// 返回为空则说明没有获取到了 bucket
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panicf("create blockBucket failed, err:%v\n", err)
			}
			// 创建一个创世块，并作为第一个区块添加到区块链中
			genesisBlock := GenesisBlock(address)
			fmt.Printf("genesisBlock:%s\n", genesisBlock)
			// 写数据, Hash 作为 key, block 作为字节流(区块的数据)
			bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
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
func GenesisBlock(address string) *Block {
	// address：为挖矿地址
	coinbase := NewCoinbaseTX(address, "Golang——自娱创世快")
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

// AddBlock 5. 添加区块
func (b *BlockChain)AddBlock(txs []*Transaction) {
	// 获取最后一个区块作为当前区块的前区块哈希
	db := b.db         // 区块链数据库
	lastHash := b.tail // 获取最后一个区块的哈希值

	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("打开 bucket 出错，不应该为空请检查")
		}
		// a. 创建新的区块
		block := NewBlock(txs, lastHash)
		// b. 添加到区块链db中
		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte("LastHashKey"), block.Hash)

		// c. 更新内存中的区块链，即将指向最后区块的区块链 tail 更新一下
		b.tail = block.Hash
		return nil
	})
}

// PrintChain 反向打印区块链
func (b *BlockChain)PrintChain() {
	blockHeight := 0
	b.db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		bucket := tx.Bucket([]byte("blockBucket"))
		// 从第一个 key--> value 进行遍历，到最后一个固定的 key 时直接返回
		bucket.ForEach(func(k, v []byte) error {
			if bytes.Equal(k, []byte("LastHashKey")) {
				return nil
			}
			block := DeSerialize(v)
			fmt.Printf("================== 区块高度：%d ====================\n", blockHeight)
			blockHeight++
			fmt.Printf("版本号：%d\n", block.Version)
			fmt.Printf("前区块哈希值：%x\n", block.PrevHash)
			fmt.Printf("梅克尔根：%x\n", block.MerkelRoot)
			fmt.Printf("时间戳：%d\n", block.TimeStamp)
			fmt.Printf("难度值（随便写的）：%d\n", block.Difficulty)
			fmt.Printf("随机数：%d\n", block.Nonce)
			fmt.Printf("当前区块哈希值：%x\n", block.Hash)
			fmt.Printf("区块数据：%s\n", block.Transactions[0].TXInputs[0].Sig)
			return nil
		})
		return nil
	})
}

// FindUTXOs 找到指定地址的所有 utxo(未消费的输出)
func (b *BlockChain)FindUTXOs(address string) []TXOutput {
	var UTXO []TXOutput
	//TODO
	return UTXO
}