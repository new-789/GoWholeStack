package main

import "fmt"

func main0601() {
	count := 0 // 用来统计循环的次数
	// 第一次循环代表公鸡，且公鸡 100 钱最多只能买 20 只，所以循环的条件小于 33
	for cock := 0; cock <= 20; cock++ {
		// 第二次循环代表母鸡，且母鸡 100 钱最多只能买 33 只，所以循环的条件为小于 33
		for hen := 0; hen <= 33; hen++ {
			// 第三次循环代表小鸡，小鸡一钱三只最多可以买 300 只，但题目要求最多只能买100只鸡，所以循环的条件为小于 100
			for chicken := 0; chicken <= 100; chicken += 1 {
				// 判断买公鸡的钱加上买母鸡的钱和买小鸡的钱等于100钱，并且公鸡，小鸡，母鸡的数量等于 100 时，则为得到的各种鸡可以买到的数量
				if cock*5+hen*3+chicken/3 == 100 && cock+hen+chicken == 100 && chicken%3 == 0 {
					fmt.Printf("公鸡：%d, 母鸡: %d, 小鸡: %d\n", cock, hen, chicken)
				}
				count++
			}
		}
	}
	fmt.Printf("循环了 %d 次\n", count)
}

func main0602() {
	count := 0 // 用来统计循环的次数
	// 第一次循环代表公鸡，且公鸡 100 钱最多只能买 20 只，所以循环的条件小于 33
	for cock := 0; cock <= 20; cock++ {
		// 第二次循环代表母鸡，且母鸡 100 钱最多只能买 33 只，所以循环的条件为小于 33
		for hen := 0; hen <= 33; hen++ {
			// 第三次循环代表小鸡，小鸡一钱三只最多可以买 300 只，但题目要求最多只能买100只鸡，所以循环的条件为小于 100
			for chicken := 0; chicken <= 100; chicken += 3 {
				// 判断买公鸡的钱加上买母鸡的钱和买小鸡的钱等于100钱，并且公鸡，小鸡，母鸡的数量等于 100 时，则为得到的各种鸡可以买到的数量
				if cock*5+hen*3+chicken/3 == 100 && cock+hen+chicken == 100 {
					fmt.Printf("公鸡：%d, 母鸡: %d, 小鸡: %d\n", cock, hen, chicken)
				}
				count++
			}
		}
	}
	fmt.Printf("循环了 %d 次\n", count)
}

func main0603() {
	chicken := 0 // 初始化小鸡的数量
	count := 0   // 同来统计循环此处
	for cock := 0; cock <= 20; cock++ {
		for hen := 0; hen <= 33; hen++ {
			chicken = 100 - (cock + hen)
			if cock*5+hen*3+chicken/3 == 100 && chicken%3 == 0 {
				fmt.Printf("公鸡：%d, 母鸡: %d, 小鸡: %d\n", cock, hen, chicken)
			}
			count++
		}
	}
	fmt.Printf("循环了 %d 次\n", count)
}
