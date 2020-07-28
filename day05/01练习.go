package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 通过键盘输入20个字母，统计个数
func main0101() {
	var arr [20]byte // 声明一个 byte 类型的数组，用来存储用户输入的数据,该数组中存储的为字母对应的 ascll 码值
	fmt.Print("Please input >>:")
	for i := 0; i < len(arr); i++ {
		fmt.Scanf("%c", &arr[i]) // 获取用户输入的内容
	}
	// 统计个数
	var ch [26]int // 用来统计字符的个数，即将字符出现的此处存储在在数组中，设定长度为 26 表示存储的内容从 'a'~~'z' 26个字母下标从 0~25
	// 记录字母出现的次数
	for i := 0; i < len(arr); i++ {
		/*  arr[i]-'a' 表示当前 arr 数组中存储的字母对应的 ascll 值减去第一个字母 ‘a’ 对应的 acsll 值得到的差值，
		该差值即为 ch 数组中的下标表示当前字母存储的位置，该位置上的默认值为 0
		++ 则表示该字母出现一次则自增一次
		*/
		ch[arr[i]-'a']++
	}

	// 打印字符出现的次数
	for i := 0; i < len(ch); i++ {
		/*
			ch[i] > 0 过滤没有出现的字母，即出现的字母才会大于零
		*/
		if ch[i] > 0 {
			// 'a'+i, 字母 a 对应的 ascll 值加上当前 i 的值即为对应的 ascll 表中对应的字母，必须使用 %c 打印，ch[i] 即为出现的次数
			fmt.Printf("字母：%c 出现 %d 次\n", 'a'+i, ch[i])
		}
	}
}

// 随机生成一注双色球彩票信息，红球6个，1~33 之间不能重复，篮球1个,取值范围 1~16 可与红球重复
func main0102() {
	// 获取随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成随机红球，范围在 1~33 之间，
	var RedBall [6]int // 声明变量用来存放生成的随机红球
	for i := 0; i < len(RedBall); i++ {
		/*
			遍历之前存在的值和新随机数是否有重复，有则进行去重操作
		*/
		temp := rand.Intn(33) + 1 // +1 表示随机数取值范围为 1~33 否则只为 0~32 之间
		for j := 0; j < i; j++ {
			// 判断当前存在的值和新的随机数是否一致，一致则生成一个新的随机数，并重置 j 的值使其再次执行该循环
			if temp == RedBall[j] {
				temp = rand.Intn(33) + 1
				j = -1 // 此处必须为 -1 否则不是从头开始遍历会导致第一个数组和其它数出现一个重复
				continue
			}
		}
		RedBall[i] = temp
	}

	// 随机生成篮球，取值范围 1~16 可与红球重复
	fmt.Println(RedBall, "+", rand.Intn(16)+1)
}
