package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main0701() {
	var num int                // 声明一个变量用来获取用户输入的数据
	inputNum := make([]int, 3) // 声明一个切片变量，用来存储用户输入的数据
	RandNum := make([]int, 3)  // 声明一个切片变量，用来存储生成的随机数数据
	// 随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成 100-999 随机数
	//random := rand.Intn(900) + 100
	// 将生成的随机数切分保存在切片
	//RandNum[0] = random / 100
	//RandNum[1] = random / 10 % 10
	//RandNum[2] = random % 10

	// 生成 100-999 随机数, 并将生成的随机数切分保存在切片
	RandNum[0] = rand.Intn(9) + 1 // 生成 1-9 随机数，该语法可防止第一位数为0
	RandNum[1] = rand.Intn(10)
	RandNum[2] = rand.Intn(10)

	var flag int = 0 // 标志位，用来确定用户是否输入正确
	// 用户输入数据
	for {
		for {
			fmt.Println("请输入三位数")
			fmt.Scan(&num)
			if num >= 100 && num <= 999 {
				break
			}
			fmt.Println("输入错误，请重新输入")
		}

		// 将用户输入法三位随机数进行切分并存在切片中
		inputNum[0] = num / 100
		inputNum[1] = num / 10 % 10
		inputNum[2] = num % 10

		// 通过循环，对两个切片中的数据做对比
		for i := 0; i < len(RandNum); i++ {
			if inputNum[i] > RandNum[i] {
				fmt.Printf("您输入的第 %d 为太大了\n", i+1)
			} else if inputNum[i] < RandNum[i] {
				fmt.Printf("您输入的第 %d 为太小了\n", i+1)
			} else {
				fmt.Printf("恭喜您猜对了，您输入的第 %d 位数相同\n", i+1)
				flag++ // 正确一位数则加一
			}
		}
		if flag == 3 {
			break
		} else {
			flag = 0
		}
	}
}
