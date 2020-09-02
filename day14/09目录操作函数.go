package main

import (
	"fmt"
	"os"
)

func main0901() {
	// 获取用户输入的目录路径
	fmt.Println("请在下方输入需要浏览的目录路径>>：")
	var path string
	fmt.Scan(&path)

	// 打开目录
	dir, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("open catalog faissled:", err)
	}
	defer dir.Close()

	// 读取目录项，-1 读取目录中的所有目录项
	fileInof, err1 := dir.Readdir(-1)
	if err1 != nil {
		fmt.Println("read catalog failed:", err1)
		return
	}

	// 遍历返回的切片
	for _, DirOrFile := range fileInof {
		// 判断是否为目录
		if DirOrFile.IsDir() {
			fmt.Printf("%s --> 是一个目录\n", DirOrFile.Name())
		} else {
			fmt.Printf("%s ==> 是一个文件\n", DirOrFile.Name())
		}
	}
}
