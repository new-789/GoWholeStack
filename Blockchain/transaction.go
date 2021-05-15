package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

// 比特币交易 Demo

// Transaction 1. 定义交易结构
type Transaction struct {
	TXID []byte  // 交易 ID
	TXInputs []TXInput  // 交易输入数组
	TXOutputs []TXOutput
}

// TXInput 定义交易结构
type TXInput struct {
	// 引用的交易ID
	TXid []byte
	// 引用的 output 的索引值
	Index int64
	// 解锁脚本，我们用地址模拟
	Sig string
}

// TXOutput 定义交易输出结构
type TXOutput struct {
	// 转账金额
	Value float64
	// 锁定脚本,我们用地址模拟
	PukKeyHash string
}

// SetHash 设置交易ID，直接将交易 Transaction 结构体进行hash 作为交易 ID
func (tx *Transaction)SetHash() {
	var buffer bytes.Buffer
	encode := gob.NewEncoder(&buffer)
	err := encode.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	data := buffer.Bytes()
	hash := sha256.Sum256(data)
	tx.TXID = hash[:]
}

// 2. 提供创建交易方法
// 3. 创建挖矿交易
// 4. 根据交易调整程序
