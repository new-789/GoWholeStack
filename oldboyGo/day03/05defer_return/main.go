package main

import "fmt"

// defer 经典案列：
// Go 语言中的函数的 return 不是原子操作，在底层是分为两步执行的：
// 1. 返回值赋值
// 2. 执行 ret 返回
// 函数中如果存在 defer 那么 defer 执行的实际是在 第一步和第二部之间
func f1() int {
	x := 5
	defer func() {
		x++ // 修改的是 x 不是返回值
	}()
	return x
}

func f2() (x int) { // x 是一个公共返回值变量
	defer func() {
		x++
	}()
	return 5 // 返回值=x
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 返回值 = y = x = 5
}

func f4() (x int) {
	defer func(x int) {
		x++ // 改变是的函数中 x 的副本
	}(x)
	return 5 // 返回值 = x = 5
}

func f5() (x int) {
	defer func(x int) int {
		x++
		return x
	}(x)
	return 5
}

// 传一个指针 x 到匿名函数
func f6() (x int) {
	defer func(x *int) {
		(*x)++
	}(&x)
	return 5 // 返回值 = x = 5  2、defer    3. RET指令
}

func main() {
	fmt.Println("f1:", f1())  // 5
	fmt.Println("f2: ", f2()) // 6
	fmt.Println("f3: ", f3()) // 5
	fmt.Println("f4: ", f4()) // 5
	fmt.Println("f5: ", f5()) // 5
	fmt.Println("f6: ", f6()) // 6
}
