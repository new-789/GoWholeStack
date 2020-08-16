package main

import (
	"fmt"
	"io"
	"os"
)

// WriteString 方法介绍
func main0601() {
	// file == var file *File
	// os.Create 创建文件时，文件不存在则创建新文件，存在则清空文件中的源内容
	file, err := os.Create("D:/test.txt")
	if err != nil {
		fmt.Println("file create fail")
		return
	}
	defer file.Close()

	// 写入文件, \n 在文件中表示换行，但是在 windows 文本文件中换行是 \r\n，在 Go 语言中一个汉字占用 3 个字符的长度，这是Go为了在 windows 系统和 Linux 系统中统一处理汉字而做的转码规则
	n, _ := file.WriteString("锄禾日当午，\r\n汗滴禾下土，\r\n谁知盘中餐，\r\n粒粒皆辛苦。")
	fmt.Println(n)
	//n, _ = file.WriteString("Hello World")
	fmt.Println(n)
}

// Write 方法使用
func main0602() {
	// 创建文件操作
	file, err := os.Create("D:/test.txt")
	if err != nil {
		fmt.Println("file create fail")
		return
	}
	// 关闭文件
	defer file.Close()

	// 将字符切片通过 Write 写入文件操作
	//b := []byte{'H', 'e', 'l', 'l', 'o'}
	//n, _ := file.Write(b)
	//fmt.Println(n)

	// 将字符串转成字节切片通过 Write 写入文件中
	str := "HelloWorld，锄禾日当午"
	//b := []byte(str)
	//n, _ := file.Write(b)
	n, _ := file.Write([]byte(str))
	fmt.Println(n)
}

// WriteAt 方法使用介绍
func main0603() {
	// 打开文件，os.OpenFile(文件名,打开模式, 打开权限)
	file, err := os.OpenFile("D:/test.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("file create fail")
		return
	}

	// 关闭文件
	defer file.Close()

	// 写入文件
	//b := []byte("Hello World, 汗滴禾下土")  // 通过字符串获取字节切片
	//// 当使用 WriteAt 进行指定位置插入数据时会依次覆盖源内容
	//n, _ := file.WriteAt(b, 0)
	//fmt.Println(n)

	// 获取文件字符个数，用作 WriteAt 方法中的偏移量信息
	//num, er := file.Seek(0, io.SeekEnd)  // 表示从文件中内容的末尾向右偏移0(也就是不偏移)个字符开始写入新的内容
	num, er := file.Seek(12, io.SeekStart) // 表示从起始位置往后偏移9个字符
	//num, er := file.Seek(-3, io.SeekEnd)  // 表示从文件中内容的末尾往左偏移三个字符开始写入新的内容
	if er != nil {
		fmt.Println("光标获取错误", er)
		return
	}

	// 写入文件操作
	b := []byte("锄禾日当午,汗滴禾下土,谁知盘中餐,粒粒皆辛苦!")
	file.WriteAt(b, num)
}
