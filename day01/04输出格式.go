package main

import "fmt"

func main0401() {
	// 输出格式 Println 打印时自带换行
	//fmt.Println("hello")
	//fmt.Println(10)
	//fmt.Println(3.14)

	// 输出格式 Print 打印时不带换行
	//fmt.Print("hello")
	//fmt.Print(10)
	//fmt.Print(3.14)

	//
	a := 10
	b := 3.1415926
	/* <font color=blue size=3>%d</font> 占位符  表示输出一个整型数据，
	**``%Nd``** 表示需要打印的内容如果不足 N 位则使用空格补足 N 位,也可在 N 前用 0 表示不足的用 0 补足
	 */
	// <font color=blue size=3>\n</font> 表示转义字符，相当于换行符
	fmt.Printf("%06d\n", a)
	/*
		<font color=blue size=3>%f</font> 占位符  表示输出一个浮点型数据，默认保留六位小数，
		指定保留小数点位数使用 **``%.Nf``** ，**``.``** 表示小数点，**``N``** 表示需要保留小数点后的位数,会对N后面的内容四舍五入
	*/
	fmt.Printf("%.3f\n", b)
}

func main0402() {
	// bool布尔  string字符串  byte字符
	var a bool // 声明 bool 类型变量，默认值为 false
	// <font color=blue size=3>%t</font> 占位符  表示输出一个布尔类型数据
	fmt.Printf("%t\n", a)

	var b string = "我的祖国"
	// <font color=blue size=3>%s</font> 占位符  表示输出一个字符串类型数据
	fmt.Printf("%s\n", b)

	var c byte = 'a'
	fmt.Println(c) // 该方法打印字符类型的数据输出的为该字符对应的 ASCII 码值
	// <font color=blue size=3>%c</font> 占位符  表示输出一个字符类型数据
	fmt.Printf("%c\n", c)
}
