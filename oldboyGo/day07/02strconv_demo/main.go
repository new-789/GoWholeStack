package main

import (
	"fmt"
	"strconv"
)

// strconv

func main() {
	// 从字符串中解析除对应的数据
	str := "10000"
	retInt64, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		fmt.Println("ParseInt failed, err:", err)
		return
	}
	fmt.Printf("type:%T， value:%d\n", retInt64, retInt64)

	// Atoi：将字符串转换成Int 类型
	retInt, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("strconv.Atoi failed, err:", err)
		return
	}
	fmt.Printf("type:%T， value:%d\n", retInt, retInt)

	// Itoa: 将数字转换为字符串数据
	retStr := strconv.Itoa(retInt)
	fmt.Println(retStr)

	// 从字符串中解析出布尔值
	boolStr := "true"
	boolValue, err := strconv.ParseBool(boolStr)
	if err != nil {
		fmt.Println("ParseBool failed ,err :", err)
		return
	}
	fmt.Printf("type:%T， value:%v\n", boolValue, boolValue)

	// 从字符串中解析除浮点数
	floatStr := "3.1415926"
	retFloat64, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("ParseFloat failed, err:", err)
		return
	}
	fmt.Printf("type:%T， value:%f\n", retFloat64, retFloat64)

	// 将数字转换成字符串类型
	i := int32(98)
	// ret := string(i)
	// fmt.Println(ret)
	ret2 := fmt.Sprintf("%d", i) // "98"
	fmt.Printf("%#v\n", ret2)
}
