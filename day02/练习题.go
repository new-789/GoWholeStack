package main

import (
	"fmt"
)

func main1() {
	var countDay int = 46      // 需要求的总天数
	var weekDay int = 7        // 一周等于 7 天
	week := countDay / weekDay // 求有多少周
	day := countDay % weekDay  // 求完周之后求余几天
	fmt.Printf("%d周%d天\n", week, day)
}

func main2() {
	seond := 107653
	// 秒数除以60即为多少分钟，多少分钟除以一小时60分即为多少小时，多少小时除以一天24小时即为当前给出秒数为多少小时，当前的小时数取余一年的 365 天即得到当前给出的秒数等于多少天
	day := seond / 60 / 60 / 24 % 365
	hour := seond / 60 / 60 % 24
	min := seond / 60 % 60
	second := seond % 60
	fmt.Printf("%d天%d时%d分%d秒", day, hour, min, second)
}

func main3() {
	var year int
	fmt.Printf("please input year>>:")
	// 获取用户输入的数据
	fmt.Scanf("%d", &year)
	// 判断是否为闰年
	b := (year%400 == 0) || (year%4 == 0 && year%100 != 0)
	fmt.Printf("%t\n", b)
}
