package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 打开文件写内容
// 使用 os 包写入文件操作
func writedemo1() {
	fileObj, err := os.OpenFile("./xxx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}

	// Write 写入文件操作
	fileObj.Write([]byte("三藏奉命去西天取经\n"))
	// WriteSTring 写入文件操作
	fileObj.WriteString("后面跟着三徒弟，悟空、悟净和悟能~!\n")
	fileObj.Close()
}

// 使用 bufio 包写文件操作
func writedemo2() {
	file, err := os.OpenFile("./xxx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	// 写入文件操作
	wr := bufio.NewWriter(file)
	// 此步骤只是将内容写在缓存中
	wr.WriteString("- 悟空会七十二变\n- 悟能会睡觉\n- 悟净会挑担~!\n")
	wr.Flush() // 将缓存的内容写入到文件
}

// 使用 ioutil 包写入文件操作
func writedemo3() {
	err := ioutil.WriteFile("./xxx.txt", []byte("路上碰到了很多妖怪"), 0755)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

func main() {
	writedemo1()
	writedemo2()
	writedemo3()
}
