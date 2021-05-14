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
	addBlock --data DATA 	"add to blockChain"
	printChan  				"print all blockChain data"
`

// Run 接收参数的动作函数
func (c *Cli)Run() {
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
		fmt.Printf("添加区块\n")
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
		fmt.Println("打印区块")
		c.PrintBlockChain()
	default:
		fmt.Println("命令无效，请检查")
		fmt.Printf(Usage)
	}
}

