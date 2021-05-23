package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
)

// 椭圆曲线算法生成公钥私钥 demo

func main() {
	// ecdsa 生成公钥和私钥
	curve := elliptic.P256() // 创建曲线
	// 生成私钥
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	// 使用私钥生成公钥
	publicKey := privateKey.PublicKey
	// 使用私钥对数据进行签名
	data := "hello world"
	hash := sha256.Sum256([]byte(data)) // 对数据进行 hash
	// 数字签名
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		log.Panic(err)
	}
	/*
	fmt.Printf("pubkey:%v\n", publicKey)
	fmt.Printf("r:%v, len:%d\n", r.Bytes(), len(r.Bytes()))
	fmt.Printf("s:%v, len:%d\n", s.Bytes(), len(s.Bytes()))
	 */

	// 将 r、s 进行拼接并序列化传输,r 和 s 的长度是一样的，后面需进行拆分
	signature := append(r.Bytes(), s.Bytes()...)
	// ... 假设这是网络传输将 signature 传给另一端

	// 经过漫长的等待假设网络另一端收到了收据，需做如下操作
	// 1. 定义两个辅助的 bit.Int
	var r1, s1 big.Int
	// 2. 拆分接收到的数据 signature，平均分，前半部分给 r，后半部分给 s
	r1.SetBytes(signature[:len(signature)/2])
	s1.SetBytes(signature[len(signature)/2:])

	// 签名校验需要的三个内容：数据、签名、公钥
	res := ecdsa.Verify(&publicKey, hash[:], &r1, &s1)
	fmt.Printf("校验结果：%v\n", res)
}
