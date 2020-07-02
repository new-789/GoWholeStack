package main

import "fmt"

func main0601() {
	// bool 类型
	var a bool
	fmt.Println(a)
	a = true
	fmt.Println(a)
	b := true
	fmt.Println(b)
}

func main0602() {
	var a float32
	var b float64
	// float32 默认小数位数保留7位有效数据，会在末尾 +1 操作
	a = 3.141592641112222211112
	// float64 默认小数位数保留i15 位又死奥数据
	b = 3.141592623443545645645
	fmt.Println(a)
	fmt.Println(b)
}

func main0603() {
	// 定义字符类型的变量
	var a byte = '0'
	var b byte = 'a'
	// byte 本身就是 uint8 类型，可以直接计算
	// 将小写字母转换成大写字母
	c := b - 32
	// 打印字符型对应的 ASCII 码对应的值
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	// 使用格式化输出 %c 可打印字符本身
	fmt.Printf("%c, %c, %c\n", a, b, c)
}

func main0604() {
	// 字符串类型
	var a string // 默认内容 "\0" 表示字符串的结束标识，如给其赋值为 "hello" 那么完整内容应该为 "hello\0", \0 默认不可见
	fmt.Println(a)

	var b string = "澳门葡京娱乐"
	var c string = "性感荷官在线发牌"
	// 字符串的相加, 将两个字符串拼接起来
	d := b + c
	fmt.Println(d)

	var e string = "字符串运算操作"
	var f string = "字符串运算操作"
	// == 运算符， 比较两个字符串内容是否相同
	fmt.Println(e == f)
}
