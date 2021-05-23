package main

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"io/ioutil"
	"log"
)

// 钱包管理功能实现demo

// Wallets 用来保存所有的 wallet 以及它的地址
type Wallets struct {
	//map[地址]钱包
	WalletsMap map[string]*Wallet
}

// NewWallets 创建方法
func NewWallets() *Wallets {
	var wallets Wallets
	wallets.WalletsMap = make(map[string]*Wallet)
	//wallets := LoadFile()
	return &wallets
}

func (w *Wallets)CreateWallet() string {
	// 获取钱包和钱包地址
	wallet := NewWallet()
	address := wallet.NewAddress()

	// 将钱包与地址添加进 map 中，key 为钱包地址，value 为 wallet 结构体
	w.WalletsMap[address] = wallet
	w.SaveToFile()
	return address
}

// SaveToFile 保存方法，将新建的 wallet 添加进去
func (w *Wallets)SaveToFile() {
	var buffer bytes.Buffer
	// gob 包对椭圆曲线计算得到的结果进行编码时，应使用 gob 的 Register 对其进行注册一下否则会报如下错误
	// panic: gob: type not registered for interface: elliptic.p256Curve
	gob.Register(elliptic.P256())  // gob 注册 elliptic.P256()
	encode := gob.NewEncoder(&buffer)
	err := encode.Encode(w)
	if err != nil {
		log.Panic(err)
	}
	if err := ioutil.WriteFile("wallet.dat", buffer.Bytes(), 0600); err != nil {
		log.Printf("将钱包地址存入文件失败：%v\n", err)
	}
}

// 读取文件方法，将所有的 wallet 读出来
