package main

import (
	"fmt"
	"os"
)

// 区块链命令行管理工具实现

type Cli struct {
	bc *BlockChain
}

const Usage = `
	addBlock --data DATA 	"添加区块链"
	printChan  				"正向打印区块链信息"
	printChainR				"反向打印区块链信息"
	getBalance --address ADDRESS "获取指定地址的余额"
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
	case "addBlock":
		// 3. 添加区块
		fmt.Printf("添加区块......\n")
		// 确保命令有效
		if len(args) == 4 && args[2] == "--data" {
			// a. 获取命令行数据
			data := args[3]
			// b. 使用 bc 的 AddBlock 方法添加数据
			c.AddBlock(data)
		} else {
			fmt.Printf("添加区块参数使用不正确，请检查！\n")
			fmt.Printf(Usage)
		}
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
	default:
		fmt.Println("命令无效，请检查")
		fmt.Printf(Usage)
	}
}
