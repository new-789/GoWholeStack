// 导入主函数的包
package main

// 导入所需的包语法, fmt==>format 标准输入输出格式包
import (
	"fmt"
)

// 双斜杠表示此行为注释内容，注释内容不参与程序编译，帮助我们理解程序的逻辑
// main 叫主函数，是程序的主入口，程序有且只有一个主函数
func main01() {
	// 在屏幕打印 hello world
	fmt.Println("hello world~!")
	fmt.Println("这是 Go 语言全栈开发学习第一个程序")
}
