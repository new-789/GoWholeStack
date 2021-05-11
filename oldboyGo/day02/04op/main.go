package main

import "fmt"

// 运算符
func main() {
	var (
		a = 5
		b = 2
	)
	// 算数运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ // 单独的语句，不能放在等号的右边  ==> a = a + 1
	b-- // 单独的语句，不能放在等号的右边  ==> b = a - 1
	fmt.Println(a)

	// 关系运算符
	fmt.Println(a == b) //Go 语言是强类型语言，相同类型的变量才能比较
	fmt.Println(a != b) // 不等于
	fmt.Println(a >= b) // 大于等于
	fmt.Println(a <= b) // 小于等于
	fmt.Println(a < b)  // 小于
	fmt.Println(a > b)  // 大于

	// 逻辑运算符
	// 如果年龄大于18岁并且年龄小于60岁
	age := 22
	if age > 18 && age < 60 {
		fmt.Println("苦逼的上班族")
	} else {
		fmt.Println("享受生活阶段")
	}
	// 如果年龄小于 18 岁，或者年龄大于 60 岁
	if age < 18 || age > 60 {
		fmt.Println("享受生活阶段")
	} else {
		fmt.Println("苦逼的上班族")
	}
	// not 取反，原本为真就为假，原本为假就为真
	isMarried := false
	fmt.Println(!isMarried) // 为真
	fmt.Println(isMarried)  // 为假

	// 位运算符，针对二进制数
	// 5 的二进制表示：101
	// 2 的二进制表示：10

	// & 按位与，两位均为1 就为 1
	fmt.Println(5 & 2) // 000
	// | 按位或：两位有一个为1就为1
	fmt.Println(5 | 2) // 111
	// ^ 按位异或：两位不一样则为 1
	fmt.Println(5 ^ 2) // 111
	// << 左移：将二进制位左移指定位数
	fmt.Println(5 << 1)  // 1010  => 10
	fmt.Println(1 << 10) // 10000000000 == >1024
	// >> 右移：将二进制位右移指定的位数
	fmt.Println(5 >> 1)  // 10 ==> 2
	var m = int8(1)      // 只能存 8 位
	fmt.Println(m << 10) // 向左移10位，没意义

	// 赋值运算符，用来给变量赋值
	// var x int
	// x = 10
	// x += 1 // x = x + 1
	// x -= 1 // x = x - 1
	// x *= 2 // x = x * 2
	// x /= 2 // x = x / 2
	// x %= 2 // x = x % 2
	// x << =2 // x = x << 2
	// x &= 2 // x = x & 2
	// x |= 3 // x = x | 3
	// x ^= 4 // x = x ^ 4
	// x >>= 2 // x = x >> 2
	// fmt.Println(x)
}
