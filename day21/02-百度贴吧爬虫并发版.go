package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func httpGet1(url string, num int) (result string, err error) {
	res, Gerr := http.Get(url)
	if Gerr != nil {
		err = Gerr
		return
	}
	defer res.Body.Close()

	buf := make([]byte, 4096)
	for {
		n, Rerr := res.Body.Read(buf)
		if n == 0 {
			fmt.Printf("=========正在爬取第>{%d}<页=======\n", num)
			break
		}
		if Rerr != nil && Rerr != io.EOF {
			err = Rerr
			return
		}
		result += string(buf[:n])
	}
	return
}

func SpiderPage(num int, page chan int) {
	url := "https://www.pianku.tv/mv/------" + strconv.Itoa(num) + ".html"
	path := "E:/CodingFiles/GolangCode/test/"

	result, err := httpGet1(url, num)
	if err != nil {
		fmt.Println("httpPage1 error:", err)
		return
	}

	// 将会去到的网页内容存入文件
	file, Cerr := os.Create(path + "第" + strconv.Itoa(num) + "页" + ".html")
	if Cerr != nil {
		fmt.Println("os.Create error:", Cerr)
		return
	}
	_, Werr := file.WriteString(result)
	if Werr != nil {
		fmt.Println("file.WriteString error:", Werr)
		return
	}

	file.Close()
	page <- num
}

func Working(start, end int) {
	page := make(chan int)
	for i := start; i <= end; i++ {
		go SpiderPage(i, page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("获取到的第 %d 网页内容存储到文件成功\n", <-page)
	}
}

func main0201() {
	var start, end int
	fmt.Print("please input start page(start >= 1):")
	fmt.Scan(&start)
	fmt.Print("please input end page(end > start):")
	fmt.Scan(&end)

	Working(start, end)
}
