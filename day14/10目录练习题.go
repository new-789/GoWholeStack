package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// 练习一：指定目录检索特定文件，从用户给出的目录中找出所有的 **.jpg** 文件
func selectJpg(file *os.File, path string) {
	// 读取目录中的内容
	dirSlice, err := file.Readdir(-1)
	if err != nil {
		fmt.Println("读取目录错误......")
	}

	// 循环遍历目录
	for _, dir := range dirSlice {
		if !dir.IsDir() {
			if strings.HasSuffix(dir.Name(), ".ipynb") {
				//fmt.Printf("%s --> 是后缀为 .jpg 格式的图片文件\n", dir.Name())
				// 调用拷贝方法将找到的文件拷贝到另一个目录下
				copyMp4(path+"/"+dir.Name(), "C:/Users/zhufeng/Desktop/test/"+dir.Name())
			}
		}
	}
}

// 练习二：指定目录拷贝特定文件，从用户给出的目录中，拷贝 **.mp3** 文件到指定目录中
func copyMp4(src, dst string) {
	openFile, er := os.Open(src)
	if er != nil {
		fmt.Println("open file err:", er)
		return
	}
	defer openFile.Close()

	// 创建需要保存的文件名
	createFile, err := os.Create(dst)
	if err != nil {
		fmt.Println("create file failed:", err)
		return
	}
	defer createFile.Close()

	// 创建读取文件的缓冲
	buf := make([]byte, 1024)
	// 循环读取文件
	for {
		n, err := openFile.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println("file read end......")
			return
		}

		// 写入文件操作
		_, er := createFile.Write(buf[:n])
		if er != nil { // 写文件一般很少出错，可以忽略该检查
			fmt.Println("save file failed:", er)
			return
		}
	}
}

func main1001() {
	fmt.Println("Please input dir>>:")
	var path string
	fmt.Scan(&path)

	// 打开目录操作
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("open file failed:", err)
		return
	}
	defer file.Close()

	selectJpg(file, path)
}
