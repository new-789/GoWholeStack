package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)
const reward = 12.5 // 挖矿奖励

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
	PubKeyHash string
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

// NewCoinbaseTX 2. 提供创建交易方法(挖矿交易)
func NewCoinbaseTX(address string, data string) *Transaction {
	// 挖矿交易的特点：
	// 1. 只有一个 input
	// 2. 无需引用交易 id
	// 3. 无需引用index
	// 旷工由于挖矿成功后金钱由系统奖励获得无需指定引用上一个节点的输出签名，
	// 所以这个 sig 字段可以由旷工自由填写数据，一般填写矿池的名字
	input := TXInput{[]byte{}, -1, data}
	// 旷工挖矿后金钱由系统的奖励获得
	output := TXOutput{reward, address}
	// 对于挖矿交易来说，只有一个input和一个output
	tx := &Transaction{[]byte{}, []TXInput{input}, []TXOutput{output}}
	tx.SetHash()
	return tx
}

// 3. 创建挖矿交易
// 4. 根据交易调整程序
