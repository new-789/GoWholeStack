package main

import (
	"fmt"
	"os"
	"strconv"
)

// 区块链命令行管理工具实现

type Cli struct {
	bc *BlockChain
}

const Usage = `
	printChan  				"正向打印区块链信息"
	printChainR				"反向打印区块链信息"
	getBalance --address ADDRESS "获取指定地址的余额"
	send FROM TO AMOUNT MINER DATA "由  from 转 amount 给 to，由 miner 挖矿，同时写入 data"
	newWallet 				"创建一个新的钱包(私钥公钥对)"
`

// Run 接收参数的动作函数
func (c *Cli) Run() {
	// 1. 得到输入的所有命令
	args := os.Args
	if len(args) < 2 {
		fmt.Printf(Usage)
		return
	}
	// 2. 分析命令
	cmd := args[1]
	// 3. 执行相应动作
	switch cmd {
	case "printChain":
		// 打印区块
		fmt.Println("打印区块......")
		c.PrintBLockChain()
	case "printChainR":
		fmt.Println("反向打印区块链.....")
		c.PrintBlockChainReverse()
	case "getBalance":
		fmt.Println("获取余额......")
		if len(args) == 4 && args[2] == "--address" {
			address := args[3]
			c.GetBalance(address)
		}
	case "send":
		fmt.Println("转账开始...........")
		if len(args) != 7 {
			fmt.Println("参数个数错误，请检查！...")
			fmt.Printf(Usage)
			return
		}
		//./block send FROM TO AMOUNT MINER DATA "由  from 转 amount 给 to，由 miner 挖矿，同时写入 data"
		from := args[2]
		to := args[3]
		amount,_ := strconv.ParseFloat(args[4], 64)  // 知识点，请注意
		miner := args[5]
		data := args[6]
		c.Send(from, to,amount, miner, data)
	case "newWallet":
		fmt.Println("创建新的钱包!....")
		c.NewWalletCommand()
	default:
		fmt.Println("命令无效，请检查")
		fmt.Printf(Usage)
	}
}
