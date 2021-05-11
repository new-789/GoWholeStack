package main

import (
	"fmt"
	"time"
)

// 时间包 time

func f1() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	// 时间戳
	fmt.Println(now.Unix())     // 秒时间戳
	fmt.Println(now.UnixNano()) // 纳秒时间戳
	// time.Unix
	timeObj := time.Unix(1616071594, 0)
	fmt.Println(timeObj.Year())
	fmt.Println(timeObj.Day())
	// 时间间隔
	fmt.Println(time.Second)
	// now + 1小时
	fmt.Println(now.Add(24 * time.Hour))
	// Sub：两个时间相减
	nextTime := now.Add(time.Hour)
	nextTime = nextTime.UTC()
	d := nextTime.Sub(time.Now())
	fmt.Println("=============", d)
	// 定时器
	// timer := time.Tick(time.Second)
	// for t := range timer {
	// 	fmt.Println(t) // 1秒钟执行一次
	// }

	// 时间格式化, 将语言中的时间对象转换成字符串类型的时间
	// 2021-03-18 21:05
	fmt.Println(now.Format("2006-01-02 15:04"))
	// 2021/03/18 21:06:15 PM ---> 12小时格式
	fmt.Println(now.Format("2006/01/02 15:04:05 PM"))
	// 2021/03/18 21:06:15.233 ---> 获取毫秒格式
	fmt.Println(now.Format("2006/03/04 15:04:05.000"))

	// 按照对应的格式解析字符串时间
	timeObj, err := time.Parse("2006-01-02", "2021-03-18")
	if err != nil {
		fmt.Println("parse time failed, err:", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())

	// Sleep：接收一个 Duration 类型变量，该变量为 int64 类型的别名
	n := 5 // int 类型
	fmt.Println("开始 sleep 了")
	// time.Duration(n) 强转为 int64 类型，等同于 int64(n)
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("5m 钟过去了")
}

// 时区
func f2() {
	now := time.Now() // 获取本地时间
	fmt.Println(now)
	// 明天的这个时间
	// 按照指定格式解析一个字符串格式的时间，该方法解析出来的时间为 UTC
	time.Parse("2006-01-02 15:04:05", "2021-03-20 20:17:30")

	// 按照东八区的时区和格式去解析一个字符串格式的时间
	loc, err := time.LoadLocation("Asia/Shanghai") // 根据字符串加载一个时区
	if err != nil {
		fmt.Println("locdLocation time failed, err:", err)
		return
	}
	fmt.Println(loc)
	// 按照指定时区解析时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-03-20 20:17:30", loc)
	if err != nil {
		fmt.Println("parse time failed, err:", err)
		return
	}
	fmt.Println(timeObj)
	// 时间对象相减
	dt := timeObj.Sub(now)
	fmt.Println(dt)
}

func main() {
	// f1
	f2()
}
