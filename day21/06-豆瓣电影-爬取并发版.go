package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

// 爬取指定 url 的页面，返回 result、error
func HttpGetDouban(url string) (result string, err error) {
	res, Gerr := http.Get(url)
	if Gerr != nil {
		err = Gerr
		return
	}
	defer res.Body.Close()

	// 循环获取整页数据
	buf := make([]byte, 4096)
	for {
		n, Rerr := res.Body.Read(buf)
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

func SavToFile(idx int, fileName, filmScore, peopleNum [][]string) {
	path := "E:/CodingFiles/GolangCode/test/" + "第" + strconv.Itoa(idx) + "页.txt"
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("os.Create error:", err)
		return
	}
	defer file.Close()

	n := len(fileName) // 得到条目数
	// 先写入抬头：电影名、演员信息、评分人数
	file.WriteString("电影名称" + "\t\t\t" + "评分" + "\t\t\t" + "评分人数" + "\r\n")
	for i := 0; i < n; i++ {
		file.WriteString(fileName[i][1] + "\t\t\t" + filmScore[i][1] + "\t\t\t" + peopleNum[i][1] + "\r\n")
	}
}

// 爬取一个豆瓣页面数据信息
func spiderDbPage(idx int, page chan int) {
	// 获取 url地址
	URL := "https://movie.douban.com/top250?start=" + strconv.Itoa((idx-1)*25) + "&filter="

	// 封装 HttpGetDouban 函数爬取 URL 对应页面
	result, err := HttpGetDouban(URL)
	if err != nil {
		fmt.Println("HttpGetDouban error:", err)
		return
	}

	// 开始横向爬取页面中具体条目的内容
	reg1 := regexp.MustCompile(`<img width="100" alt="(.*?)"`) // 解析编译正则表达式---返回电影名
	// 提取需要的信息
	fileName := reg1.FindAllStringSubmatch(result, -1) // 提取内容

	// 获取评分分数
	reg2 := regexp.MustCompile(`<span class="rating_num" property="v:average">(?s:(.*?))</span>`)
	filmScore := reg2.FindAllStringSubmatch(result, -1)

	// 获取评分人数
	reg3 := regexp.MustCompile(`<span>\d+人评价</span>`)
	peopleNum := reg3.FindAllStringSubmatch(result, -1)

	// 保存有用信息内容到文件
	SavToFile(idx, fileName, filmScore, peopleNum)
	// 与主 go 程配合完成同步
	page <- idx
}

func toWork(start, end int) {
	fmt.Printf("我正在爬取 %d 到 %d 页......\n", start, end)
	page := make(chan int) // 防止主 go 程提前结束
	for i := start; i <= end; i++ {
		go spiderDbPage(i, page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第 %d 页爬取完毕\n", <-page)
	}
}

func main0601() {
	// 指定爬取的起始页、终止页
	var start, end int

	fmt.Print("Please input start Page(>=1):")
	fmt.Scan(&start)
	fmt.Print("Please input end Page(>=start):")
	fmt.Scan(&end)

	toWork(start, end)
}
