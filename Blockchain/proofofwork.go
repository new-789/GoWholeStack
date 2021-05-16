package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// 工作量证明 Demo

// ProofOfWork 工作量证明结构
type ProofOfWork struct {
	// block
	block *Block
	// 目标值,
	// bit.Int: 一个非常大的数，它有丰富的方法：比较，赋值等方法
	target *big.Int
}

// NewProofOfWork POW 工作量证明函数
func NewProofOfWork(block *Block) *ProofOfWork {
	// 指定的难度值，是一个 string 类型，需要进行转换
	targetStr := "0000100000000000000000000000000000000000000000000000000000000000"
	// 临时变量，目的是将上面的 string 类型的难度值转成 big.Int 类型
	tmpInt := big.Int{}
	// 将难度值赋值给 big.Int，指定16进制的格式
	tmpInt.SetString(targetStr, 16)
	pow := &ProofOfWork{
		block: block,
		target: &tmpInt,
	}
	return pow
}

// Run 根据随机数不断执行哈希计算找到符合条件的哈希值
func (p *ProofOfWork)Run() ([]byte, uint64) {
	var nonce uint64
	b := p.block
	var hash [32]byte
	fmt.Println("开始挖矿。。。。。。。")
	for {
		// 1. 拼装数据（区块的数据，还有不断变化的随机数）
		tmp := [][]byte{
			Uint64ToByte(b.Version),
			b.PrevHash,
			b.MerkelRoot,
			Uint64ToByte(b.TimeStamp),
			Uint64ToByte(b.Difficulty),
			Uint64ToByte(nonce),
			// MerkelRoot 根只对区块头做 hash 值，区块体通过 MerkelRoot 产生影响
			//b.Data,
		}
		// 将二维切片数字通过一维切片拼接起来，返回一个一维的切片
		blockInfo := bytes.Join(tmp, []byte(""))
		// 2. 做哈希运算
		hash = sha256.Sum256(blockInfo)
		// 3. 与 pow 中的 target 进行比较
		tmpInt := big.Int{}
		// 将我们得到的哈希数组转换成一个 big.Int 类型
		tmpInt.SetBytes(hash[:])
		// 比较当前的哈希值与目标哈希值，如果当前的哈希值小于目标哈希值说明找到了，否则继续找

		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		//
		//func (x *Int) Cmp(y *Int) (r int) {
		if tmpInt.Cmp(p.target) == -1 {  // big.Int 包中的数值比较方法
			// a. 找到了，退出返回
			fmt.Printf("挖矿成功, hash:%x, nonce: %d\n", hash, nonce)
			return hash[:], nonce
		} else {
			// b. 没找到，继续找，随机数加1
			nonce++
		}
	}
	//return hash[:], nonce
}
