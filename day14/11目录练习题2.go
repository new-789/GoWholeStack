package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// 打开目录读取目录内容函数
func OpenDir(path string) {
	// 打开目录操作
	dirInfo, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("打开目录错误，错误为：", err)
		return
	}
	defer dirInfo.Close()

	dirNames := make([]string, 0)
	// 读取目录内容操作
	if names, err := dirInfo.Readdir(-1); err != nil {
		fmt.Println("获取目录中的文件内容失败,错误信息：", err)
		return
	} else {
		for _, name := range names {
			// 判断如果不是目录才对其进行操作，减少程序的负担
			if !name.IsDir() {
				if strings.HasSuffix(name.Name(), ".txt") {
					// 将所有为 .txt 后缀的文件加入到切片，便于后边循环遍历读取文件内容
					dirNames = append(dirNames, name.Name())
				}
			}
		}
		// 打开文件并开始按行读取文件中的内容，传入参数为文件的路径和存储文件名的切片
		ReadInfo(path+"/", dirNames)
	}
}

// 打开文件读取文件内容函数
func ReadInfo(src string, dirNames []string) {
	// 创建切片用来存储分割后所有的文件内容
	nameSlice := make([]string, 0)
	// 循环打开文件
	for _, name := range dirNames {
		file, err := os.OpenFile(src+name, os.O_RDONLY, 6)
		if err != nil {
			fmt.Println("打开文件错误，错误信息为:", err)
			return
		}
		defer file.Close()

		// 读取文件内容操作
		reader := bufio.NewReader(file)
		for {
			buf, err := reader.ReadBytes('\n')
			if err != nil && err == io.EOF {
				fmt.Println("文件读取完成......")
				break
			} else if err != nil {
				fmt.Println("readBytes err:", err)
				break
			}
			nameSlice = append(nameSlice, strings.Fields(string(buf[:]))...)
		}
		// 调用函数统计文件中单词的个数
		CountNum(nameSlice)
	}
}

// 统计单词出现的个数
func CountNum(fileInfo []string) {
	// 创建 map 用来存放单词出现的次数
	count := make(map[string]int)
	// 循环统计单词个数，并存放置在 map 中
	for i := 0; i < len(fileInfo); i++ {
		if _, has := count[fileInfo[i]]; !has {
			count[fileInfo[i]] = 1
		} else {
			count[fileInfo[i]] = count[fileInfo[i]] + 1
		}
	}
	// 此处打印的内容可以优化为参数以防止输出的内容写死
	fmt.Printf("目录所有文件中 love 单词共有 %d 个 \n", count["love"])
	return
}

func main11() {
	// 接收用户输入的目录路径
	fmt.Println("请输入要找寻的目录>>:")
	var path string
	fmt.Scan(&path)
	// 调用打开目录函数开始执行
	OpenDir(path)
}
