package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kataras/iris/v12"
)

//
func main3005() {
	app := iris.New()
	// 一、通过程序代码对应用进行全局配置
	app.Configure(iris.WithConfiguration(iris.Configuration{
		// 如果设置为 true，当人为中断程序执行时，则不会自动正常将服务器关闭，如果设置为 true 需要自己自定义处理
		DisableInterruptHandler: false,
		// 该配置项表示更正并将请求的路径重定向到已注册的路径，如，如果请求 /home/ 但找不到此 router 的处理程序，然后路由检查 /home 处理程序是否存在，如果是，(permant) 将客户端重定向到正常，默认Wie false
		DisablePathCorrection: false,
		//
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "Mon,02 Jan 2006 15:04:05 GMT",
		Charset:                           "utf-8",
	}))
	// 二、通过读取 tml 配置文件读取服务配置，**注意：要在 run 方法运行之前执行**
	app.Configure(iris.WithConfiguration(iris.TOML("/home/zhufeng/CodingFiles/GolangCode/src/github.com/GoWholeStack/day30 Iris/config/iris.tml")))
	// 三、通过读取 yaml 配置文件读取服务配置，同样要在 run 方法之前执行
	app.Configure(iris.WithConfiguration(iris.YAML("/home/zhufeng/CodingFiles/GolangCode/src/github.com/GoWholeStack/day30 Iris/config/iris.yml")))

	// 四、通过 json 配置文件进行应用配置
	file, err := os.Open("/home/zhufeng/CodingFiles/GolangCode/src/github.com/GoWholeStack/day30 Iris/config/iris.json")
	if err != nil {
		fmt.Println("打开配置文件错误", err)
		return
	}
	defer file.Close()
	// json 解码器
	decoder := json.NewDecoder(file)
	conf := Configuration{}
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Println("/home/zhufeng/CodingFiles/GolangCode/src/github.com/GoWholeStack/day30 Iris/config", err)
		return
	}
	fmt.Println(conf.Port)
	app.Run(iris.Addr(":8080"))
}

type Configuration struct {
	AppName string `json:"appname"`
	Port    int    `json:"port"`
}
