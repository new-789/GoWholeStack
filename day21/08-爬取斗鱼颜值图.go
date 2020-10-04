package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func SaveImg(url string, idx int, page chan int) {
	path := "E:/CodingFiles/GolangCode/test/" + strconv.Itoa(idx+1) + ".jpg"
	f, ferr := os.Create(path)
	if ferr != nil {
		fmt.Println("os.Create error:", ferr)
		return
	}
	defer f.Close()

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("saveImg error:", err)
		return
	}
	defer res.Body.Close()

	buf := make([]byte, 4096)
	for {
		n, err := res.Body.Read(buf)
		if n == 0 {
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println("saveImg in http.get error:", err)
			return
		}
		f.Write(buf[:n])
	}
	page <- idx
}

func main() {
	url := "https://www.douyu.com/g_yz"

	result, err := httpget(url)
	if err != nil {
		fmt.Println("httpget error:", err)
		return
	}

	reg := regexp.MustCompile(`<img loading="lazy" src="(.*?)"`)
	alls := reg.FindAllStringSubmatch(result, -1)

	page := make(chan int)
	for idx, imgAddr := range alls {
		go SaveImg(imgAddr[1], idx, page)
	}

	n := len(alls)
	for i := 0; i < n; i++ {
		fmt.Printf("下载第 %d 张图片完成\n", <-page) // 防止主 go 程退出
	}
}

func httpget(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return
}
