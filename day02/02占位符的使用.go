package main

import "fmt"

func main0201() {
	//var a int64 = 10
	a := 10               // int
	fmt.Printf("%d\n", a) // 打印整型数据

	//var b float64 = 10
	//b := 10 // 使用该语法自动推导式定义的 b := 10 为 int，类型不能使用 %f 格式化输出
	b := 10.0             // float64
	fmt.Printf("%f\n", b) // 打印浮点型数据

	var c bool = true
	fmt.Printf("%t\n", c) // 打印bool 类型的值

	var d byte = 'A'      // byte == uint8
	fmt.Printf("%c\n", d) // 打印 byte 类型的值

	var e string = "hello china"
	fmt.Printf("%s\n", e) // 打印字符串类型的数据

	fmt.Printf("%p\n", &a) // 取变量 a 的内存地址并输出

	// %T 打印变量对应的数据类型
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)

	// %% 打印一个百分号字面量
	fmt.Printf("35%%")
}

func main0202() {
	// 计算机能够识别的进制 二进制、八进制、十进制、十六进制
	a := 123   // 十进制数据
	b := 0123  // 八进制数据 以 0 开头的数据是八进制数据
	c := 0xabc // 十六进制数据 以 0x 开头的数据是十六进制
	// GO 语言中不能直接表示二进制数据

	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)

	// %b 占位符 表示输出一个二进制数据
	//fmt.Printf("a 的二进制值为：%b\n", a)
	//fmt.Printf("b 的二进制值为：%b\n", b)
	//fmt.Printf("c 的二进制值为：%b\n", c)

	// %o 占位符 表示输出一个八进制数据
	//fmt.Printf("a 的八进制值为: %o\n", a)
	//fmt.Printf("b 的八进制值为: %o\n", b)
	//fmt.Printf("c 的八进制值为: %o\n", c)

	// %x 占位符  表示输出十六进制数据, a-f 为小写
	fmt.Printf("a 的十六进制值为: %x\n", a)
	fmt.Printf("b 的十六进制值为: %x\n", b)
	fmt.Printf("c 的十六进制值为: %x\n", c)

	// %X 占位符  表示输出十六进制数据, a-f 为大写
	fmt.Printf("a 的十六进制值为: %X\n", a)
	fmt.Printf("b 的十六进制值为: %X\n", b)
	fmt.Printf("c 的十六进制值为: %X\n", c)
}
