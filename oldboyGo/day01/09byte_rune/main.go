package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "hello 沙河"
	// len() 求得的是 byte 字节的数量
	fmt.Println(len(s))

	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(s[i])
	// 	fmt.Printf("%c\n", s[i]) // 占位符 %c ，表示打印字符
	// }

	// for _, i := range s { // 从字符串中拿出具体的字符
	// 	fmt.Printf("%c\n", i)
	// }

	// 字符串修改
	s2 := "白龙马"
	s3 := []rune(s2) // 将字符串强制转换成了一个 rune 切片
	s3[0] = '红'
	fmt.Println(string(s3)) // 将 rune 切片强制转换成字符串

	c1 := "红"
	c2 := '红'
	fmt.Printf("c1:%T c2:%T\n", c1, c2)

	c3 := "H"
	c4 := 'H'
	fmt.Printf("c3:%T c4:%T\n", c3, c4)

	// // 类型转换
	n := 10
	var f float64
	f = float64(n)
	fmt.Printf("%T,%v\n", f, f)

	ss := "hello沙河小王子"
	floag := 0
	for _, v := range ss {
		if unicode.Is(unicode.Han, v) {
			floag++
		}
	}
	fmt.Println(floag)
}
