package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// net/http client

/*
声明全局变量创造一个关闭了长连接的 client 适用于请求比较频繁的场景，使其保持始终只用这一个 client
var (
	client = http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false
		},
	}
)
*/

func main() {
	// 向服务器发起 get 请求
	/* 第一种方式：发送简单的 GET 请求
	res, err := http.Get("http://127.0.0.1:8081/test/?name=三藏&age=999")
	if err != nil {
		fmt.Println("http get failed, err:", err)
		return
	}
	 */

	// 第二种方式：构建一个请求对象，及所带的参数
	// 1. url values 用来给 url 中定义请求时需要带的参数
	data := url.Values{}
	// 2. 解析请求 url 地址
	urlObj,err := url.ParseRequestURI("http://127.0.0.1:8081/test/")
	if err != nil {
		fmt.Println("解析 url 错误：",err)
		return
	}
	// 3. 设置请求需要带的参数
	data.Set("name", "三藏")
	data.Set("age", "1000")
	// 4. url encode 对参数字符进行编译之后的 URL
	urlStr := data.Encode()
	fmt.Println(urlStr)
	// 5. 将参数放入解析好的 urlObj 中
	urlObj.RawQuery = urlStr
	// 6. 创建一个请求对象,包含请求方法、字符串类型的请求对象(urlObj.String(),将请求对象转换为字符串类型)
	req, err := http.NewRequest("GET", urlObj.String(),nil)
	if err != nil {
		fmt.Println("newRequest failed, err:", err)
		return
	}
	// 7. 发请求
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("defaultClient failed, err:", err)
		return
	}
	defer res.Body.Close() // 请求发送完成之后一定要关闭网络 IO

	/* 第三种方式：禁用 KeepAlive(持久连接) 的 client, 适用于请求不是特别频繁，用完即关闭该连接
	tr := &http.Transport{
		DisableKeepAlives: false,
	}
	client := http.Client{
		Transport: tr,
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("disableKeepAlive request failed, err:", err)
		return
	}
	defer res.Body.Close()
	 */

	// 从 res 中读取服务端返回的数据
	// var data []byte
	//res.Body.Read()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read res.Body failed, err:", err)
		return
	}
	fmt.Println(string(b))
}
