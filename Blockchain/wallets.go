package main

// 钱包管理功能实现demo

// Wallets 用来保存所有的 wallet 以及它的地址
type Wallets struct {
	//map[地址]钱包
	WalletsMap map[string]*Wallet
}

// NewWallets 创建方法
func NewWallets() *Wallets {
	// 获取钱包和钱包地址
	wallet := NewWallet()
	address := wallet.NewAddress()
	// 将钱包与地址添加进 map 中，key 为钱包地址，value 为 wallet 结构体
	var wallets Wallets
	wallets.WalletsMap = make(map[string]*Wallet)
	wallets.WalletsMap[address] = wallet
	return &wallets
}

// 保存方法，将新建的 wallet 添加进去

// 读取文件方法，将所有的 wallet 读出来
