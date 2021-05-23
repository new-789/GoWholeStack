package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
	"log"
)

// 钱包, 这里的钱包是一个结构，每一个钱包保存了公钥、私钥对

// Wallet 钱包结构体
type Wallet struct {
	Private *ecdsa.PrivateKey
	// 约定，这里的 pubKey 不存储原始的公钥，而是存储 X和Y拼接的字符串，在校验段重新拆分(参考ecdsa demo 中的 r、s 传递)
	PubKey []byte
}

// NewWallet 创建钱包
func NewWallet() *Wallet {
	// 创建曲线
	curve := elliptic.P256()
	// 生成私钥
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	// 使用私钥生成公钥
	publicKeyOrig := privateKey.PublicKey
	// 拼接公钥中的 x，y (x/y 为椭圆曲线图中的坐标)
	pubKey := append(publicKeyOrig.X.Bytes(), publicKeyOrig.Y.Bytes()...)
	return &Wallet{Private: privateKey, PubKey: pubKey}
}

// NewAddress 生成地址
func (w *Wallet)NewAddress() string {
	// 对公钥做 sh256 和 rip160 计算
	pubKey := w.PubKey
	rip160HashValue := HashPuKey(pubKey)
	version := byte(00) // 版本号
	// 对 rip160 得到的结果与 version 进行拼接
	paylood := append([]byte{version}, rip160HashValue...)
	// checksum
	checkCode := CheckSum(paylood)
	// 将获取到的四个字节的校验码数据与 paylood 进行拼接得到 25 字节数据
	paylood = append(paylood, checkCode...)

	// go 语言有一个叫做 btcd 的库，这是 go 语言实现的比特币全节点
	// base58 包下载地址: github.com/btcsuite/btcutil
	address := base58.Encode(paylood)
	return address
}

// HashPuKey 对公钥进行 hash 及 rip160 计算
func HashPuKey(data []byte) []byte {
	// 对公钥进行 sha256 计算
	hash := sha256.Sum256(data)
	// 理解为生成 hash160 编码器，通过 ripemd160 对公钥的 sha256 计算结果再次进行 rip160 计算
	rip160hasher := ripemd160.New()
	// 将需要进行 rip160 计算的数据添加进来
	_, err := rip160hasher.Write(hash[:])
	if err != nil {
		log.Panic(err)
	}
	// 进行 rip160 计算,返回 rip160 的哈希结果
	rip160HashValue := rip160hasher.Sum(nil)
	return rip160HashValue
}

// CheckSum 对需要生成的校验码数据进行两次 sha256 计算，并取前四个字节校验码数据
func CheckSum(data []byte) []byte {
	// checksum
	// 两次 sha256
	hash1 := sha256.Sum256(data)
	hash2 := sha256.Sum256(hash1[:])
	// 前四字节校验码
	checkCode := hash2[:4]
	return checkCode
}