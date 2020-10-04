package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func HttpGet(url string, num int) (result string, err error) {
	res, err1 := http.Get(url)
	if err1 != nil {
		err = err1 // 将封装函数内容的错误返回给调用者
		return
	}
	defer res.Body.Close()

	// 循环读取网页数据，传出给调用者
	buf := make([]byte, 4096)
	for {
		n, err2 := res.Body.Read(buf)
		if n == 0 {
			fmt.Printf("爬取第 %d 网页完成\n", num)
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		// 累加每次循环读到的 buf 数据，存入 result，一次性返回
		result += string(buf[:])
	}
	return
}

//爬取页面操作
func working(start, end int) {
	fmt.Printf("正在爬取 %d 页到 %d 页......\n", start, end)
	// 循环爬取每一页数据
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		result, err := HttpGet(url, i)
		if err != nil {
			fmt.Println("HttpGet error:", err)
			return
		}
		// 将爬取到的整页数据存储到文件中
		path := "E:/CodingFiles/GolangCode/test/"
		file, Rerr := os.Create(path + "第 " + strconv.Itoa(i) + " 页" + ".html")
		if Rerr != nil {
			fmt.Println("os.Create error:", Rerr)
			return
		}

		_, Werr := file.WriteString(result)
		if Werr != nil {
			fmt.Println("file.Write error:", Werr)
			return
		}
		// 保存好一个文件就关闭一个文件
		file.Close()
	}
}

func main0101() {
	//指定爬取起始页和终止页面
	var start, end int
	fmt.Print("请输入爬取的起始页(>=1): ")
	fmt.Scan(&start)
	fmt.Print("请输入爬取的终止页(>=start): ")
	fmt.Scan(&end)

	working(start, end)
}
