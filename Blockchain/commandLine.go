package main

import (
	"fmt"
)

// 命令行逻辑功能实现Demo

// PrintBLockChain 正向打印区块链信息
func (c *Cli) PrintBLockChain() {
	c.bc.PrintChain()
	fmt.Printf("打印区块链完成")
}

// PrintBlockChainReverse 反向打印区块链信息
func (c *Cli) PrintBlockChainReverse() {
	// 调用迭代器，返回每一个区块数据
	iterator := c.bc.NewIterator() // 创建一个迭代器

	for {
		// 返回区块并左移
		block := iterator.Next()
		for _, tx := range block.Transactions {
			fmt.Println(tx)
		}
		/*
		fmt.Printf("====================================\n")
		fmt.Printf("版本号：%d\n", block.Version)
		fmt.Printf("前一区块哈希值：%x\n", block.PrevHash)
		fmt.Printf("梅克尔根：%x\n", block.MerkelRoot)
		timeFormat := time.Unix(int64(block.TimeStamp), 0).Format("2006-01-02 15:04:05")
		fmt.Printf("时间戳：%s\n", timeFormat)
		fmt.Printf("难度值(随便写的后面改进)：%d\n", block.Difficulty)
		fmt.Printf("随机数：%d\n", block.Nonce)
		fmt.Printf("当前区块哈希值：%x\n", block.Hash)
		fmt.Printf("区块数据：%s\n", block.Transactions[0].TXInputs[0].PubKey)
		 */
		// 判断如果单钱区块链的前一个区块切片长度为零则说明迭代完毕，则退出
		if len(block.PrevHash) == 0 {
			fmt.Printf("区块链遍历结束！")
			break
		}

	}
}

func (c *Cli) GetBalance(address string) {
	// 1. 校验地址
	if !IsValidAddress(address) {
		fmt.Println("地址无效", address)
		return
	}
	// 2. 生成公钥哈希
	pubKeyHash := GetPubKeyFromAddress(address)
	utxos := c.bc.FindUTXOs(pubKeyHash)
	total := 0.0
	for _, utxo := range utxos {
		total += utxo.Value
	}
	fmt.Printf("\"%s\" 的余额为：%f\n", address, total)
}

func (c *Cli) Send(from, to string, amount float64, miner, data string) {
	fmt.Printf("from:%s to:%s amount:%f miner:%s data:%s\n", from, to, amount, miner, data)
	// 校验地址
	if !IsValidAddress(from) {
		fmt.Println("地址无效 from", from)
		return
	}
	if !IsValidAddress(to) {
		fmt.Println("地址无效 to", to)
		return
	}
	if !IsValidAddress(miner) {
		fmt.Println("地址无效 miner", miner)
		return
	}

	// 1. 创建挖矿交易
	conbase := NewCoinbaseTX(miner, data)
	// 2. 创建一个普通交易
	tx := NewTransaction(from, to, amount, c.bc)
	if tx == nil {
		fmt.Println("无效的交易")
		return
	}
	// 3. 添加到区块
	c.bc.AddBlock([]*Transaction{conbase, tx})
	fmt.Println("转账成功")
}

// NewWalletCommand 创建一个新钱包地址
func (c *Cli) NewWalletCommand() {
	ws := NewWallets()
	address := ws.CreateWallet()
	fmt.Printf("wallet地址:%s\n", address)
}

// GetListAddresses 获取所有钱包地址并打印
func (c *Cli) GetListAddresses() {
	ws := NewWallets()
	addresses := ws.GetAllAddresses()
	for _, address := range addresses {
		fmt.Printf("wallet 地址:%v\n", address)
	}
}
