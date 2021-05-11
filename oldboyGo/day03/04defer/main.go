package main

import "fmt"

// defer 语句：多用于函数结束之前释放资源(文件句柄、数据库连接、socket)
func deferDemo() {
	fmt.Println("start")
	// 将 defer 后面的语句延迟执行，延迟到函数即将返回时运行
	// 一个函数中可以有多个 defer，多个 defer 语句按照后进先出原则延迟执行
	defer fmt.Println("嘿嘿嘿")
	defer fmt.Println("呵呵呵")
	defer fmt.Println("哈哈哈")
	fmt.Println("end")
}

func main() {
	deferDemo()
}
