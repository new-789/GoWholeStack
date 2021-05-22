package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"
	"log"
	"time"
)

// block 区块业务代码

//Block 1. 定义区块链结构
type Block struct {
	// 1 版本号
	Version uint64
	// 2 前区块哈希
	PrevHash []byte
	// 3 Merkel 根（梅克尔根，这就是一个哈希值，先不管，后面介绍）
	MerkelRoot []byte
	// 4 时间戳
	TimeStamp uint64
	// 5 难度值
	Difficulty uint64
	// 6 随机数，以为就是挖矿要找的数据
	Nonce uint64

	// a 当前区块哈希，正常比特币区块中没有当前区块的哈希，为了实现方便做了简化
	Hash []byte
	// b 数据
	//Data []byte
	// 真实的交易数组
	Transactions []*Transaction
}

//1. 补充区块字段
//2. 重新极端哈希函数
//3. 优化代码

// Uint64ToByte 辅助函数，将 uint64 转成 []byte
func Uint64ToByte(num uint64) []byte  {
	// 将 uint64 转换为 []byte 类型
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}

//NewBlock 2. 创建区块
func NewBlock(txs []*Transaction, precHash []byte) *Block {
	block := &Block{
		Version: 00,
		PrevHash: precHash,
		MerkelRoot: []byte{},
		TimeStamp: uint64(time.Now().Unix()),
		Difficulty: 0,  // 随便填写的无效值
		Nonce: 0,  // 同上
		Hash:     []byte{}, // 先填空，后面计算
		//Data:     []byte(data),
		Transactions: txs,
	}
	block.MerkelRoot = block.MakeMerkelRoot()
	// 创建一个 pow 对象
	pow := NewProofOfWork(block)
	// 查找模目标的随机数，不停的进行哈希运算
	hash, nonce := pow.Run()
	// 根据挖矿结果对区块数据进行更新
	block.Hash = hash
	block.Nonce = nonce
	return block
}

// Serialize 序列化
func (b *Block)Serialize() []byte {
	// 编码后的数据放到 buf 中
	var buffer bytes.Buffer
	//使用 gob 进行序列化(编码)得到字节流
	// 1. 定义一个编码器
	encoder := gob.NewEncoder(&buffer)
	// 2. 使用编码器对结构体进行编码
	err := encoder.Encode(&b)
	if err != nil {
		log.Panicf("编码出错:%v\n", err)
	}
	return buffer.Bytes()
}

// DeSerialize 反序列化
func DeSerialize(data []byte) *Block {
	decoder := gob.NewDecoder(bytes.NewBuffer(data))
	var block *Block
	err := decoder.Decode(&block)
	if err != nil {
		log.Panicf("解码数据失败：%v\n", err)
	}
	return block
}

/*
// SetHash 3. 生成哈希
func (b *Block)SetHash() {
	// 1. 拼装数据
	/*
	blockInfo = append(blockInfo, Uint64ToByte(b.Version)...)
	blockInfo = append(blockInfo, b.PrevHash...)
	blockInfo = append(blockInfo, b.MerkelRoot...)
	blockInfo = append(blockInfo, Uint64ToByte(b.TimeStamp)...)
	blockInfo = append(blockInfo, Uint64ToByte(b.Difficulty)...)
	blockInfo = append(blockInfo, Uint64ToByte(b.Nonce)...)
	blockInfo = append(blockInfo, b.Data...)
	*/
	/*
	// 对上面冗余代码进行优化
	tmp := [][]byte{
		Uint64ToByte(b.Version),
		b.PrevHash,
		b.MerkelRoot,
		Uint64ToByte(b.TimeStamp),
		Uint64ToByte(b.Difficulty),
		Uint64ToByte(b.Nonce),
		b.Data,
	}
	// 将二维切片数字通过一维拼接起来，返回一个一维的切片
	blockInfo := bytes.Join(tmp, []byte(""))
	// 2. sha256 生成哈希值
	hash := sha256.Sum256(blockInfo)
	b.Hash = hash[:] // 给当前区块添加哈希值
}
*/

// MakeMerkelRoot 模拟梅克尔根，只对交易的数据做简单的拼接，不做二叉树处理
func (b *Block)MakeMerkelRoot() []byte {
	var info []byte
	//var finalInfo [][]byte
	// 将交易的哈希值拼接起来在做整体做处理
	for _, tx := range b.Transactions {
		info = append(info, tx.TXID...)
		//finalInfo = [][]byte{tx.TXID}
	}
	hash := sha256.Sum256(info)
	return hash[:]
}