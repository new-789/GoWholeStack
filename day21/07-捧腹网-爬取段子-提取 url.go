package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// 获取一个网页所有的内容
func HttpGetPF(url string) (result string, err error) {
	// 请求 url 开始爬取网页内容
	resp, Gerr := http.Get(url)
	if Gerr != nil {
		err = Gerr
		return
	}
	defer resp.Body.Close()

	// 从获取到的网页内容中读取内容信息
	buf := make([]byte, 4096)
	for {
		n, Rerr := resp.Body.Read(buf)
		if n == 0 {
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

// 将爬取到的段子内容写入文件操作
func SaveToFile(titleSlice, contentSlice []string, idx int) {
	path := "E:/CodingFiles/GolangCode/test/" + "第" + strconv.Itoa(idx) + "页.txt"
	f, Ferr := os.Create(path)
	if Ferr != nil {
		fmt.Println("os.Create error:", Ferr)
		return
	}
	defer f.Close()

	// 写入文件操作
	n := len(titleSlice)
	for i := 0; i < n; i++ {
		f.WriteString("标题:" + titleSlice[i] + "\n" + contentSlice[i] + "\n")
		f.WriteString("==================================================\n")
	}
}

// 抓取一个网页，带有十个段子对应十个 URL
func SpiderPagePF(idx int, page chan int) {
	// 拼接 URL
	PageUrl := "https://m.pengfue.com/xiaohua_" + strconv.Itoa(idx) + ".html"
	// 封装函数获取段子的 url
	result, err := HttpGetPF(PageUrl)
	if err != nil {
		fmt.Println("HttpGetPF error:", err)
		return
	}
	// 解析编译正则，
	reg := regexp.MustCompile(`<h1 class="f18"><a href="(.*?)"`)
	// 获取每个段子对应的 url
	urlAddr := reg.FindAllStringSubmatch(result, -1)
	// 创建用于存储 title 和 content 的切片,用来保存获取到的整页内容
	titleSlice := make([]string, 0)
	contentSlice := make([]string, 0)
	for _, v := range urlAddr {
		// 根据获取到的 url 爬取一个页面的 title 和 content
		title, content, err := SpiderJokePage(v[1])
		if err != nil {
			fmt.Println("SpiderJokePage error:", err)
			continue
		}
		// 将 title 和 content 内容追加到切片末尾
		titleSlice = append(titleSlice, title)
		contentSlice = append(contentSlice, content)
	}
	// 将获取到的内容保存到文件中
	SaveToFile(titleSlice, contentSlice, idx)
	// 防止主 go 程提前结束
	page <- idx
}

// 爬取一个段子页面的 title 和 content
func SpiderJokePage(url string) (title, content string, err error) {
	result, Herr := HttpGetPF(url)
	if Herr != nil {
		err = Herr
		return
	}

	// 编译、解析正则表达式
	reg1 := regexp.MustCompile(`<title>(.*?)</title>`)
	// 获取 title 内容
	titleInfo := reg1.FindAllStringSubmatch(result, -1)[0][1]
	// 由于该网页获取标题最好的方式是从 title 标签中获取，但有其它内容，所以进行切分并取切片中第 0 个内容
	title = strings.Split(titleInfo, "-")[0]
	// 编译、解析正则表达式
	reg2 := regexp.MustCompile(`<div class="con-txt">(?s:(.*?))</div>`)
	// 获取段子 content 内容,并去除换行符 \n 与制表符 \t
	content = strings.Trim(strings.Trim(reg2.FindAllStringSubmatch(result, -1)[0][1], "\n"), "\t")
	return
}

func ToWork(start, end int) {
	fmt.Printf("开始爬取 %d 页到 %d 页网页中的数据............\n", start, end)
	page := make(chan int)
	for i := start; i <= end; i++ {
		// 开启 go 程并发爬取内容
		go SpiderPagePF(i, page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第 %d 页数据爬取完成......\n", <-page)
	}
}

func main0701() {
	var start, end int
	fmt.Print("请输入需要爬取页面的起始页(>=1):")
	fmt.Scan(&start)
	fmt.Print("请输入需要爬取页面的结束页(>=start):")
	fmt.Scan(&end)

	ToWork(start, end)
}
