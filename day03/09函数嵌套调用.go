package main

import "fmt"

func test2(a, b int) {
	sum := a + b
	fmt.Println(sum)
}
func test1(a, b int) {
	test2(a, b)
}

func main0901() {
	test1(10, 20)
}

func test4(args ...int) {
	for i, data := range args {
		fmt.Println("下标", i, "值", data)
	}
}

func test3(args ...int) {
	// 如果函数参数为不定参，传递方式为 **``参数名[0:]...``** 表示从不定参数下标为0的数据开始传递，一直传递到最后一个， **``...``** 表示不定参数据的格式
	//test4(args[0:]...)  等同于 test4(args...)
	//test4(args[2:]...)  从下标为2的数据开始传递参数，到最后一个结束
	//test4(args[:2]...)  从下标为0的数据开始传递参数，到第二个结束(注：不包含下标为 2 的值)
	test4(args[1:3]...) // 从下标为1的数据开始传递参数，待第三个结束，不包含下标为 3 的值
}

func main0902() {
	test3(1, 2, 3, 4, 5)
}
