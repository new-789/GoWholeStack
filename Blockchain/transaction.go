package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"math/big"
	"strings"
)

const reward = 12.5 // 挖矿奖励

// 比特币交易 Demo

// Transaction 1. 定义交易结构
type Transaction struct {
	TXID      []byte    // 交易 ID
	TXInputs  []TXInput // 交易输入数组
	TXOutputs []TXOutput
}

// TXInput 定义交易输入结构
type TXInput struct {
	// 引用的交易ID
	TXid []byte
	// 引用的 output 的索引值
	Index int64
	// 解锁脚本，我们用地址模拟
	//Sig string

	// 真正的数字签名,由 r，s 拼成的[]byte
	Signature []byte
	// 约定，这里的 pubKey 不存储原始的公钥，而是存储 X和Y拼接的字符串，在校验段重新拆分(参考ecdsa demo 中的 r、s 传递)
	PubKey []byte // 注意，是公钥不是哈希也不是地址
}

// TXOutput 定义交易输出结构
type TXOutput struct {
	// 转账金额
	Value float64
	// 锁定脚本,我们用地址模拟
	//PubKeyHash string

	// 收款方公钥的哈希，注意，是哈希而不是公钥，也不是地址
	PubKeyHash []byte
}

// NewTXOutput 给 TXOutput 提供一个创建的方法，否则无法使用 Lock
func NewTXOutput(value float64, address string) *TXOutput {
	output := TXOutput{
		Value: value,
	}
	output.Lock(address)
	return &output
}

// Lock 由于现在存储的字段是地址的公钥哈希，所以无法直接创建 TXOutput
// 为了能够得到公钥哈希，我们需要写一个 Lock 函数进行处理一下
func (output *TXOutput) Lock(address string) {
	/*
		// 1. 解码
		// 2. 截取出公钥哈希：去除开头的 version(1字节) 和结尾的校验码(4字节）
	*/
	// 真正的锁定动作！！！
	output.PubKeyHash = GetPubKeyFromAddress(address)
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
	/*
		// 1. 交易收入 Input 只有一个
		if len(tx.TXInputs) == 1 {
			input := tx.TXInputs[0]
			// 2. 交易 id 为空 || 交易的 Index 为 -1 则说明该交易为挖矿交易
			if !bytes.Equal(input.TXid, []byte{}) || input.Index != -1 {
				return false
			}
		}
	*/
	// 判断交易收入只有一个，且交易 id 等于 0，并且 交易 index 等于 -1，则说明为挖矿交易
	if len(tx.TXInputs) == 1 && len(tx.TXInputs[0].TXid) == 0 && tx.TXInputs[0].Index == -1 {
		return true
	}
	return false
}

// NewCoinbaseTX 2. 提供创建交易方法(挖矿交易)
func NewCoinbaseTX(address, data string) *Transaction {
	// 挖矿交易的特点：
	// 1. 只有一个 input
	// 2. 无需引用交易 id
	// 3. 无需引用index
	// 旷工由于挖矿成功后金钱由系统奖励获得无需指定引用上一个节点的输出签名，
	// 所以这个 PubKey 字段可以由旷工自由填写数据，一般填写矿池的名字
	// 签名先填写为空，后面创建完整交易后，最后做一次签名即可
	input := TXInput{[]byte{}, -1, nil, []byte(data)}
	// 旷工挖矿后金钱由系统的奖励获得
	//output := TXOutput{reward, address}
	// 新的创建方法
	output := NewTXOutput(reward, address)
	// 对于挖矿交易来说，只有一个input和一个output
	tx := Transaction{[]byte{}, []TXInput{input}, []TXOutput{*output}}
	tx.SetHash()
	return &tx
}

// NewTransaction 3. 创建普通转账交易
func NewTransaction(from, to string, amount float64, bc *BlockChain) *Transaction {
	// 1. 创建交易之后要进行数数字签名，--> 所以需要私钥打开钱包"NewWallets()"
	wallets := NewWallets()
	// 2. 找到自己的钱包，根据地址返回自己的 wallet
	wallet := wallets.WalletsMap[from]
	if wallet == nil {
		fmt.Printf("没有找到该地址的钱包，交易创建失败!\n")
		return nil
	}
	// 3. 得到对应的公钥、私钥
	PubKey := wallet.PubKey
	privateKey := wallet.Private

	// 得到公钥哈希
	pubKeyHash := HashPubKey(PubKey)

	// 1. 找到最合理的 UTXO 集合 map[string]uint64，和计算后转账的总金额
	utxos, resValue := bc.FindNeedUTXOs(pubKeyHash, amount)
	if resValue < amount {
		fmt.Printf("余额不足，交易失败！\n")
		return nil
	}
	var inputs []TXInput
	var outputs []TXOutput
	// 2. 创建交易输入，将找到的这些 UTXO 逐一转成 inputs
	for id, indexArray := range utxos {
		for _, i := range indexArray {
			input := TXInput{[]byte(id), int64(i), nil, PubKey}
			inputs = append(inputs, input)
		}
	}
	// 3. 创建交易输出 outputs
	//output := TXOutput{amount, to}
	output := NewTXOutput(amount, to)
	outputs = append(outputs, *output)
	// 4. 如有零钱则找零
	if resValue > amount {
		output = NewTXOutput(resValue-amount, from)
		outputs = append(outputs, *output)
	}
	tx := &Transaction{[]byte{}, inputs, outputs}
	tx.SetHash()
	// 进行签名操作
	bc.SignTransaction(tx, privateKey)
	return tx
}

// Sign 签名的具体实现，参数为：私钥、inputs 里面所有引用的交易结构 map[string]Transaction
// map[2222]Transaction222
// map[3333]Transaction333
func (tx *Transaction) Sign(privateKey *ecdsa.PrivateKey, prevTXs map[string]Transaction) {
	// 如果是挖矿讲义则不进行签名
	if tx.IsCoinbase() {
		return
	}
	// 1. 创建一个当前交易的副本：txCopy, 使用函数：TrimmedCopy：要把 Signature 和 PubKey 字段设置为 nil
	txCopy := tx.TrimmedCopy()
	// 2. 循环遍历txCopy的inputs，得到该 input 所引用的 output 的公钥哈希
	for i, input := range txCopy.TXInputs {
		prevTx := prevTXs[string(input.TXid)]
		if len(tx.TXID) == 0 {
			log.Panic("引用的交易无效....!")
		}
		// 注：不要对 input 进行赋值，这是一个副本，要对 txCopy.TXInput[xx] 进行操作，否则无法把 pubKeyHash 传进来
		txCopy.TXInputs[i].PubKey = prevTx.TXOutputs[input.Index].PubKeyHash

		// 所需要的三个数据都具备了，开始做哈希处理
		// 3. 生成要签名的数据，要签名的数据一定是哈希值
		// 3.1：我们对每一个 input 都要签名一次，签名的数据是由当前 input 引用的 output 的哈希+当前 outputs(都承载在当前这个txCopy 中)
		// 3.2：要对这个拼号的 txCopy 进行哈希处理，SetHash得到TXID，这个 TXID 就是我们要签名最终数据
		txCopy.SetHash()
		// 还原，以免影响到后面 Input 的签名
		txCopy.TXInputs[i].PubKey = nil
		signDataHash := txCopy.TXID
		// 4. 执行签名，得到 r,s 字节流
		r, s, err := ecdsa.Sign(rand.Reader, privateKey, signDataHash)
		if err != nil {
			log.Panic(err)
		}
		// 5. 放到我们所签名的 input 的 Signature 中
		signature := append(r.Bytes(), s.Bytes()...)
		tx.TXInputs[i].Signature = signature
	}
}

// TrimmedCopy 复制一个交易的副本,并修改 inputs 中的 Signature 和 PubKey 字段置为空
func (tx *Transaction) TrimmedCopy() Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	for _, input := range tx.TXInputs {
		inputs = append(inputs, TXInput{input.TXid, input.Index, nil, nil})
	}

	for _, output := range tx.TXOutputs {
		outputs = append(outputs, output)
	}
	return Transaction{tx.TXID, inputs, outputs}
}

// Verify 签名校验
	// 校验过程分析
	// 所需要的数据：公钥、数据(txCopy 生成哈希)、签名
	// 我们要对每一个签名过的 input 进行签名校验
func (tx *Transaction) Verify(prevTXs map[string]Transaction) bool {
	// 如果是挖矿交易则不用进行校验
	if tx.IsCoinbase() {
		return true
	}
	// 1. 得到签名数据
	txCopy := tx.TrimmedCopy()
	// 遍历原始数据
	for i, input := range tx.TXInputs {
		prevTx := prevTXs[string(input.TXid)]
		if len(prevTx.TXID) == 0 {
			log.Panic("引用的交易无效")
		}
		txCopy.TXInputs[i].PubKey = prevTx.TXOutputs[input.Index].PubKeyHash
		txCopy.SetHash()
		dataHash := txCopy.TXID
		// 2. 得到 signature，反推回 r,s 用来与 txCopy 中的签名进行校验
		sidnature := input.Signature //拆 r, s
		r := big.Int{}
		s := big.Int{}
		r.SetBytes(sidnature[:len(sidnature)/2])
		s.SetBytes(sidnature[len(sidnature)/2:])
		// 3. 得到PubKey，x,y 得到原生的公钥
		pubKey := input.PubKey
		x := big.Int{}
		y := big.Int{}
		x.SetBytes(pubKey[:len(pubKey)/2])
		y.SetBytes(pubKey[len(pubKey)/2:])
		pubKeyOrigin := ecdsa.PublicKey{Curve:elliptic.P256(), X:&x, Y:&y}
		// 4. Verify 校验，如果校验不通过则返回 false
		if !ecdsa.Verify(&pubKeyOrigin, dataHash, &r, &s) {
			return false
		}
	}
	return true
}

// String 格式化打印区块信息
func (tx *Transaction) String() string {
	var lines []string
	lines = append(lines, fmt.Sprintf("-------> Transaction %x:", tx.TXID))
	for i, input := range tx.TXInputs {
		lines = append(lines, fmt.Sprintf("        Input %d:", i))
		lines = append(lines, fmt.Sprintf("			TXID %x:", input.TXid))
		lines = append(lines, fmt.Sprintf("			Index %d:", input.Index))
		lines = append(lines, fmt.Sprintf("			Signature %x:", input.Signature))
		lines = append(lines, fmt.Sprintf("			PubKey %x:", input.PubKey))
	}
	for i, output := range tx.TXOutputs {
		lines = append(lines, fmt.Sprintf("        Onput %d:", i))
		lines = append(lines, fmt.Sprintf("			Value %f:", output.Value))
		lines = append(lines, fmt.Sprintf("			PubKeyHash %x:", output.PubKeyHash))
	}
	return strings.Join(lines, "\n")
}