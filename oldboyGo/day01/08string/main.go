package main

import (
	"fmt"
	"strings"
)

// 字符串
func main() {
	// 在程序中 \ 本来是具有特殊含义的，使用中应使用 \ 将反斜线转义成一个单纯的 \
	path := "\"d:\\Go\\src\\code.exe\""
	fmt.Println(path)

	s := "I'm ok"
	fmt.Println(s)

	// 多行字符串
	s2 := `
	第一行
	第二行
	第三行
	`
	fmt.Println(s2)
	s3 := `d:\Go\src\code.exe`
	fmt.Println(s3)

	// 字符串相关操作
	fmt.Println(len(s3))

	// 字符串拼接
	name := "李霞"
	world := "小仙女"
	nw := name + world
	fmt.Println(nw)
	ww := fmt.Sprintf("%s%s", name, world)
	// fmt.Printf("%s%s\n", name, world)
	fmt.Println(ww)

	// 字符串分割
	ret := strings.Split(s3, `\`)
	fmt.Println(ret)

	// 包含
	fmt.Println(strings.Contains(ww, "李霞"))
	// 前缀判断
	fmt.Println(strings.HasPrefix(ww, "李霞"))
	// 后缀判断
	fmt.Println(strings.HasSuffix(ww, "李霞"))

	s4 := "abcdeb"
	//  判断字段出现的位置
	fmt.Println(strings.Index(s4, "c"))
	// 判断子串最后一次出现的位置
	fmt.Println(strings.LastIndex(s4, "b"))

	// join 对字符串切片拼接操作
	fmt.Println(strings.Join(ret, `\`))
}
