package main

import (
	"github.com/GoWholeStack/oldboyGo/day08/mylogger"
)

var log mylogger.Logger // 声明一个全局接口变量

// 测试自定义日志库
func main() {
	log = mylogger.NewConseLog("Info")                                   // 输出日志到终端实例
	log = mylogger.NewFileLogger("Info", "./", "test.log", 10*1024*1024) // 输入日志到文件实例
	id := 10010
	name := "三藏"
	for {
		log.Debug("这是一条 debug 日志")
		log.Tarce("这是一条 Tarce 级别日志")
		log.Info("这是一条 info 日志")
		log.Warning("这是一条 warning 级别日志")
		log.Error("这是一条 error 级别日志, id:%d name:%s", id, name)
		log.Fatal("这是一条 fatal 级别日志")

		// time.Sleep(time.Second)
	}
}
