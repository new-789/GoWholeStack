package main

import (
	"fmt"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

// 写出dispatchCoin 函数实现分金币
func dispatchCoin() int {
	for _, v := range users {
		// 字符串循环出来得到的是一个个的字符
		for _, c := range v {
			// switch 判断，满足条件的分金币，并减去总金币数量
			switch c {
			case 'e', 'E':
				distribution[v]++
				coins--
			case 'i', 'I':
				distribution[v] += 2
				coins -= 2
			case 'o', 'O':
				distribution[v] += 3
				coins -= 3
			case 'u', 'U':
				distribution[v] += 4
				coins -= 4
			}
		}
	}
	for k, v := range distribution {
		fmt.Printf("%s ： %1d\n", k, v)
	}
	return coins
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}
