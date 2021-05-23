package main

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"io/ioutil"
	"log"
	"os"
)

// 钱包管理功能实现demo

const walletFileName = "wallet.dat"

// Wallets 用来保存所有的 wallet 以及它的地址
type Wallets struct {
	//map[地址]钱包
	WalletsMap map[string]*Wallet
}

// NewWallets 创建方法
func NewWallets() *Wallets {
	var wallets Wallets
	wallets.WalletsMap = make(map[string]*Wallet)
	wallets.LoadFile()
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
	if err := ioutil.WriteFile(walletFileName, buffer.Bytes(), 0600); err != nil {
		log.Printf("将钱包地址存入文件失败：%v\n", err)
	}
}

// LoadFile 读取文件方法，将所有的 wallet 读出来
func (w *Wallets)LoadFile() {
	// 读文件之前检查文件是否存在，不存在则直接退出
	_, err := os.Stat(walletFileName)
	if os.IsNotExist(err) { // 判断需要读取的文件是否存在
		return
	}
	// 读取文件内容
	content, err := ioutil.ReadFile(walletFileName)
	if err != nil {
		log.Panic(err)
	}
	gob.Register(elliptic.P256()) // gob 注册 elliptic.P256()
	// 解码操作
	decode := gob.NewDecoder(bytes.NewReader(content))
	var ws Wallets
	err = decode.Decode(&ws)
	if err != nil {
		log.Panic(err)
	}
	// 对于结构来说，里面有 map 的，需要指定来赋值，不要在最外层进行赋值
	w.WalletsMap = ws.WalletsMap
}

// GetAllAddresses 获取所有钱包地址
func (w *Wallets)GetAllAddresses() []string {
	var addresses []string
	// 遍历钱包将所有的 key 取出来保存在数据中并返回
	for address := range w.WalletsMap {
		addresses = append(addresses, address)
	}
	return addresses
}