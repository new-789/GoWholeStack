package main

import "fmt"

// 递归：自己调用自己，递归一定要有一个明确的退出条件，
// 递归适合处理那种问题相同、问题规模越来越小的场景
// 永远不要高估自己

// 3! = 3*2*1  = 3 * 2!
// 4! = 4*3*2*1  = 4 * 3!
// 5! = 5*4*3*2*1  = 5 * 4!

// 计算 n 的阶乘
func f(n int64) int64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

// 上台阶面试题：n 个台阶，一次可以走一步，也可以走两步，有多少种走法
func f1(n int64) int64 {
	if n <= 1 {
		return 1 // 如果只有一个台阶就只有一种走发
	} else if n <= 2 {
		return 2 // 如果只有两个台阶则只有两种走法
	}
	return f1(n-1) + f1(n-2)
}

func main() {
	// ret := f(7)
	// fmt.Println(ret)

	ret := f1(4)
	fmt.Println(ret)
}
