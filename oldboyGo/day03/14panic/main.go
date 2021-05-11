package main

import "fmt"

// panic 和 recover
func funcA() {
	fmt.Println("a")
}

func funcB() {
	// 刚刚打开了数据库连接
	defer func() {
		err := recover() // 尝试回复 panic 错误让程序继续执行，但程序中尽量不要用
		fmt.Println(err)
		if err != nil {
			fmt.Println("释放数据库连接......")
		}
	}()
	panic("出现了严重的错误！！！") // 程序遇到 panic 程序直接崩溃退出
	fmt.Println("b")
}

func funcC() {
	fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()
}
