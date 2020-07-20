package main

// 使用其它目录下的函数，需要导入其它目录中的包名
import "github.com/FullStackDevelStudy/day04/src/userinfo"

// 全局变量
var num int = 123

// 在同级别目录下，包名要相同
func main() {
	add(10, 20)
	// 在不同级别目录下调用其它目录下中的函数
	userinfo.Login()
	userinfo.SelectUser()
	userinfo.DeleteUser()
}
