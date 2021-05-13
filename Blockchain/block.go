package main

import (
	"bytes"
	"encoding/binary"
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
	Data []byte
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
func NewBlock(data string, precHash []byte) *Block {
	block := &Block{
		Version: 00,
		PrevHash: precHash,
		MerkelRoot: []byte{},
		TimeStamp: uint64(time.Now().Unix()),
		Difficulty: 0,  // 随便填写的无效值
		Nonce: 0,  // 同上
		Hash:     []byte{}, // 先填空，后面计算
		Data:     []byte(data),
	}
	//block.SetHash()
	// 创建一个 pow 对象
	pow := NewProofOfWork(block)
	// 查找模目标的随机数，不停的进行哈希运算
	hash, nonce := pow.Run()
	// 根据挖矿结果对区块数据进行更新
	block.Hash = hash
	block.Nonce = nonce
	return block
}

func (b *Block)toByte() []byte {
	// TODO
	return []byte{}
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