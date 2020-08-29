package main

import (
	"fmt"
	"strconv"
)

// FormatBool 使用示例
func main0201() {
	// 将其他类型转换成字符串
	s := strconv.FormatBool(false)
	fmt.Printf("转换后类型为：%T, 值为：%s", s, s)
}

// Itoa 使用示例
func main0202() {
	a := 999
	// 将整型数据转换为字符串
	s := strconv.Itoa(a)
	fmt.Printf("转换后类型为诶：%T，值为：%s", s, s)
}

// FormatInt 使用示例
func main0203() {
	var d int64 = 999
	// 将 int64 为数据按照 2 进制格式进行转换
	s := strconv.FormatInt(d, 2)
	fmt.Printf("转换后类型为诶：%T，值为：%s", s, s)
}

// FormatFloat 使用示例
func main0204() {
	var f float64 = 3.1415926
	// 将浮点类型数据转换为字符串
	s := strconv.FormatFloat(f, 'f', 3, 64)
	fmt.Printf("转换后类型为诶：%T，值为：%s", s, s)
}

// ====================================================================================================================
// ParseBool 将字符串转换为布尔使用示例
func main0205() {
	str := "false"
	b, e := strconv.ParseBool(str)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("转换后类型为诶：%T, 值为：%v", b, b)
	}
}

// Atoi 将字符串转换为整型使用示例
func main0206() {
	d := "999"
	v, err := strconv.Atoi(d)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("转换后类型为诶：%T, 值为：%v", v, v)
	}
}

// ParseInt 使用示例
func main0207() {
	s := "111110111"
	v, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("转换后类型为诶：%T, 值为：%v", v, v)
	}
}

// ParseFloat 使用示例
func main0208() {
	str := "3.1415926"
	v, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("转换后类型为诶：%T, 值为：%v", v, v)
	}
}

// ==============================================================================

// AddendInt 使用示例
func main0209() {
	b := make([]byte, 0, 1024)
	b = strconv.AppendInt(b, 999, 10)
	fmt.Println(string(b)) // 将字节类型切片转换为 string 类型并输出
}

// AppendBool 使用示例
func main0210() {
	b := make([]byte, 0, 1024)
	b = strconv.AppendBool(b, false)
	fmt.Println(string(b)) // 将字节类型切片转换为 string 类型并输出
}

// AppendFloat 使用示例
func main0211() {
	b := make([]byte, 0, 1024)
	b = strconv.AppendFloat(b, 3.1415926234234, 'f', -1, 64)
	fmt.Println(string(b))
}

// AppendQuote 使用示例
func main0212() {
	b := make([]byte, 0, 1024)
	b = strconv.AppendQuote(b, "helloWorld")
	fmt.Println(string(b))
}
