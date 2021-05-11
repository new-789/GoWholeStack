package main

import (
	"fmt"
	"io"
	"os"
)

// 文件操作

func f1() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
}

// 在文件中间插入内容
func f2() {
	// 打开要操作的文件
	file, err := os.OpenFile("./test.html", os.O_RDONLY, 0755)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	// defer file.Close()

	// 因为没有直接在文件中插入内容，所以要借助一个临时文件
	tmpFile, err := os.OpenFile("./sb.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Printf("create file failed, er:%v\n", err)
		return
	}
	// defer tmpFile.Close()
	// 读取要操作的文件，然后将读出来的内容存入临时文件中
	var ret [1]byte
	n, err := file.Read(ret[:])
	if err != nil {
		fmt.Printf("read file failed, err:%v\n", err)
		return
	}
	// 将内容写入临时文件
	tmpFile.Write(ret[:n])

	// 写入要插入的内容
	_, err = file.Seek(int64(n), 0) // 光标移动
	if err != nil {
		fmt.Printf("seek failed err:%v\n", err)
		return
	}
	_, err = tmpFile.Write([]byte{'c'})
	if err != nil {
		fmt.Printf("write file failed, err: %v\n", err)
		return
	}
	// 接着将源文件后面的内容写入临时文件
	var d [128]byte
	for {
		n, err := file.Read(d[:])
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件读取完毕~！")
				tmpFile.Write(d[:n])
				break
			}
			fmt.Printf("read2 file failed, err:%v\n", err)
			return
		}
		tmpFile.Write(d[:n])
	}
	// 源文件后续的也写入了临时文件,然后将临时文件重命名为源文件名
	file.Close()
	tmpFile.Close()
	os.Rename(tmpFile.Name(), file.Name())
}

func main() {
	f2()
}
