package main

import "math/big"

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
	// 1. 拼装数据
	// 2.
	return []byte("helloWorld"), 10
}
