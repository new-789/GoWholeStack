package main

import (
	"bytes"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"time"
)

// 区块的哈希值（即区块链的链条） Demo

// BlockChain 4. 引入区块链
type BlockChain struct {
	// 定义一个区块链数组
	db   *bolt.DB
	tail []byte // 用来存储最后一个区块的哈希值
}

const (
	blockChainDb = "blockChain.db"
	blockBucket  = "blockBucket"
)

// NewBlockChain 5. 定义一个区块链
func NewBlockChain(address string) *BlockChain {
	// 最后一个区块的哈希，从 bolt 数据库读出来的
	var lastHash []byte
	//  打开数据库
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
func (b *BlockChain) AddBlock(txs []*Transaction) {
	for _, tx := range txs {
		if !b.VerifyTransaction(tx) {
			fmt.Println("旷工发现无效交易....!")
			return
		}
	}
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
func (b *BlockChain) PrintChain() {
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
			timeOut := time.Unix(int64(block.TimeStamp), 0).Format("2006-01-02 15:04:05")
			fmt.Printf("时间戳：%s\n", timeOut)
			fmt.Printf("难度值（随便写的）：%d\n", block.Difficulty)
			fmt.Printf("随机数：%d\n", block.Nonce)
			fmt.Printf("当前区块哈希值：%x\n", block.Hash)
			fmt.Printf("区块数据：%s\n", block.Transactions[0].TXInputs[0].PubKey)
			return nil
		})
		return nil
	})
}

// FindUTXOs 找到指定地址的所有 utxo(未消费的上级区块输出)
func (b *BlockChain) FindUTXOs(pubKeyHash []byte) []TXOutput {
	var UTXO []TXOutput
	txs := b.FindUTXOTransactions(pubKeyHash)
	for _, tx := range txs {
		for _, output := range tx.TXOutputs {
			if bytes.Equal(pubKeyHash, output.PubKeyHash) {
				UTXO = append(UTXO, output)
			}
		}
	}
	return UTXO
}

// FindNeedUTXOs 根据需求找到最合理的 UTXO 集合，返回 map[string]uint64
func (b *BlockChain) FindNeedUTXOs(senderPubKeyHash []byte, amount float64) (map[string][]uint64, float64) {
	//找到的合理的 utxos 集合
	utxos := make(map[string][]uint64)
	// 找到的 UTXOS 里面包含的金额总数，即需要交易的金额
	var calc float64
	txs := b.FindUTXOTransactions(senderPubKeyHash)
	for _, tx := range txs {
		for i, output := range tx.TXOutputs {
			// 当交易中的转账人地址与转出方的地址相同时则比较转账金额，满足则直接返回 utxos,calc，不满足则继续统计
			//if from == output.PubKeyHash {
			// 两个 []byte 的比较，直接比较是否相同，返回 True 和 false
			if bytes.Equal(senderPubKeyHash, output.PubKeyHash) {
				if calc < amount {
					utxos[string(tx.TXID)] = append(utxos[string(tx.TXID)], uint64(i))
					// 统计当前 utxo 的总额
					calc += output.Value
					// 通过计算满足条件后则返回 utxos,calc
					if calc >= amount {
						return utxos, calc
					}
				} else {
					fmt.Printf("不满足转账金额，当前总额：%f， 目标金额：%f\n", calc, amount)
				}
			}
		}
	}
	return utxos, calc
}

// FindUTXOTransactions 用来遍历所有交易，寻找所有上游链条中 output 集合中跟我有关的交易加入到 Transaction 中
func (b *BlockChain) FindUTXOTransactions(senderPubKeyHash []byte) []*Transaction {
	// 用来存储所有包含 utxo 交易集合
	var txs []*Transaction
	//定义一个map来保存消费过的output，key是这个output的交易id，value是这个交易中索引的数组 map[交易id][]int64
	spentOutputs := make(map[string][]int64)

	//创建迭代器
	it := b.NewIterator()
	for {
		// 遍历区块
		block := it.Next()
		// 遍历交易
		for _, tx := range block.Transactions {
		OUTPUT:
			// 遍历 output 找到和自己相关的 utxo （在添加 output 之前检查一下是否以及消耗掉）
			for i, output := range tx.TXOutputs {
				// 在这里做一个过滤，将所有消耗过的 outputs 和当前的所即将添加的 Output 进行对比
				// 如果相同则跳过，否则添加
				// 如果当前交易id 存在于标识的 map 中，那么说明这个交易里面有消耗过的 output
				if spentOutputs[string(tx.TXID)] != nil { // 不为空则说明存在与 map 中
					for _, j := range spentOutputs[string(tx.TXID)] {
						if int64(i) == j {
							continue OUTPUT
						}
					}
				}
				//这个output和我们目标的地址相同，满足条件，加到返回UTXO数组中
				if bytes.Equal(output.PubKeyHash, senderPubKeyHash) {
					// !!!!!! 重点，返回所有包含我的 output 交易的集合
					txs = append(txs, tx)
				}
			}
			// 如果当前交易是挖矿交易，则不做遍历直接跳过
			if !tx.IsCoinbase() {
				// 遍历 inputs 找到自己花费过的 utxo 的集合(将消耗过的标示出来)
				for _, input := range tx.TXInputs {
					// 首先对公钥做哈希处理，然后进行判断如果当前 input 与目标一致，则说明该交易是消耗过的output，就加入 map 中
					pubKeyHash := HashPubKey(input.PubKey)
					if bytes.Equal(pubKeyHash, senderPubKeyHash) {
						spentOutputs[string(input.TXid)] = append(spentOutputs[string(input.TXid)], input.Index)
					}
				}
			} else {
				//fmt.Printf("这是coinbase 挖矿交易，不做 input 遍历")
			}
		}
		if len(block.PrevHash) == 0 {
			fmt.Printf("区块遍历完毕，程序退出！...\n")
			break
		}
	}
	return txs
}

// FindTransactionByTXid 根据 id 查找交易本身
func (b *BlockChain)FindTransactionByTXid(id []byte) (Transaction, error) {
	// 初始化迭代器
	it := b.NewIterator()
	for {
		// 1. 遍历区块链
		block := it.Next()
		// 2. 遍历交易
		for _, tx := range block.Transactions {
			// 3. 比较交易，找到了直接退出
			if bytes.Equal(tx.TXID, id) {
				return *tx, nil
			}
		}
		if len(block.PrevHash) == 0 {
			fmt.Println("区块链遍历结束....！")
			break
		}
	}
	// 4. 如果没找到，返回空 Transaction ，同时返回错误状态
	return Transaction{}, errors.New("交易没有找到,无效的交易id,请检查")
}

// SignTransaction 签名
func (b *BlockChain)SignTransaction(tx *Transaction, privateKey *ecdsa.PrivateKey) {
	// 签名，交易创建的最后进行签名
	prevTXs := make(map[string]Transaction)
	// 找到所有引用的交易
	// 1. 根据 inputs 来找，有多少个 input 就遍历多少次
	// 2. 找到目标的交易(根据TXid来找)
	// 3. 添加到 prevTXs 中
	for _, input := range tx.TXInputs {
		// 根据 id 查找交易本身，需要遍历整个区块链
		tx, err := b.FindTransactionByTXid(input.TXid)
		if err != nil {
			log.Panic(err)
		}
		prevTXs[string(input.TXid)]=tx
	}
	tx.Sign(privateKey, prevTXs)
}

// VerifyTransaction 签名校验
func (b *BlockChain) VerifyTransaction(tx *Transaction) bool {
	if tx.IsCoinbase() {
		return true
	}
	// 签名，交易创建的最后进行签名
	prevTXs := make(map[string]Transaction)
	// 找到所有引用的交易
	// 1. 根据 inputs 来找，有多少个 input 就遍历多少次
	// 2. 找到目标的交易(根据TXid来找)
	// 3. 添加到 prevTXs 中
	for _, input := range tx.TXInputs {
		// 根据 id 查找交易本身，需要遍历整个区块链
		tx, err := b.FindTransactionByTXid(input.TXid)
		if err != nil {
			log.Panic(err)
		}
		prevTXs[string(input.TXid)]=tx
	}
	return tx.Verify(prevTXs)
}