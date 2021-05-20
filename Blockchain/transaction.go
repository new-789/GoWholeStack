package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

const reward = 12.5 // 挖矿奖励

// 比特币交易 Demo

// Transaction 1. 定义交易结构
type Transaction struct {
	TXID      []byte    // 交易 ID
	TXInputs  []TXInput // 交易输入数组
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
func (tx *Transaction) SetHash() {
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

// IsCoinbase 函数，用来判断当前交易是否为挖矿交易
func (tx *Transaction) IsCoinbase() bool {
	// 1. 交易收入 Input 只有一个
	if len(tx.TXInputs) == 1 {
		input := tx.TXInputs[0]
		// 2. 交易 id 为空 || 交易的 Index 为 -1 则说明该交易为挖矿交易
		if !bytes.Equal(input.TXid, []byte{}) || input.Index != -1 {
			return false
		}
	}
	return true
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

// NewTransaction 3. 创建普通转账交易
func NewTransaction(from, to string, amount float64, bc *BlockChain) *Transaction {
	// 1. 找到最合理的 UTXO 集合 map[string]uint64，和计算后转账的总金额
	utxos, resValue := bc.FindNeedUTXOs(from, amount)
	if resValue < amount {
		fmt.Printf("余额不足，交易失败！\n")
		return nil
	}
	var inputs []TXInput
	var outputs []TXOutput
	// 2. 创建交易输入，将找到的这些 UTXO 逐一转成 inputs
	for id, indexArray := range utxos {
		for _, i := range indexArray {
			input := TXInput{[]byte(id), int64(i), from}
			inputs = append(inputs, input)
		}
	}
	// 3. 创建交易输出 outputs
	output := TXOutput{amount, to}
	outputs = append(outputs, output)
	// 4. 如有零钱则找零
	if resValue > amount {
		outputs = append(outputs, TXOutput{resValue-amount, from})
	}
	tx := &Transaction{[]byte{}, inputs, outputs}
	tx.SetHash()
	return tx
}

// 4. 根据交易调整程序
