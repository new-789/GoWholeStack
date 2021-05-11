package main

import "fmt"

// 数组
// 存放元素的容器，必须指定存放元素的类型和容量（长度）,
// 数组的长度是数组类型的一部分
func main() {
	var a1 [3]bool // [true false true]
	var a2 [4]bool // [true true false false]
	fmt.Printf("a1:%T a2:%T\n", a1, a2)

	// 数组的初始化,如果不初始化默认都是零值（布尔值：false，整型和浮点型：0，字符串：""）
	fmt.Println(a1, a2)
	// 初始化方式一
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	// 初始化方式二：根据初始值自动推断数组的长度是多少
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a10)
	// 初始化方式三：根据索引进行初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)

	// 数组的遍历
	citys := [...]string{"北京", "上海", "深圳"} // 索引：0~2
	// 1. 根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(i, citys[i])
	}
	// for range 遍历索引
	for index, value := range citys {
		fmt.Println(index, value)
	}

	// 多维数组 [[1 2] [3 4] [5 6]]
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)

	// 二维数组的遍历
	for _, v := range a11 {
		fmt.Println(v)
		for _, d := range v {
			fmt.Println(d)
		}
	}

	// 数组是值类型
	b1 := [3]int{1, 2, 3} // [1 2 3]
	b2 := b1              // [1 2 3] ctrl+c  ctrl+v  ==> 相当于吧一个文档从文件夹 a 拷贝到文件夹 b
	b2[0] = 100           // [100 2 3]
	fmt.Println(b1, b2)

	// 练习题
	//1. 求数组 [1,3,5,7,8] 所有元素的和。
	aa := [5]int{1, 3, 5, 7, 8}
	var result int
	for _, v := range aa {
		result += v
	}
	fmt.Println(result)

	// 2.找出数组中和为指定值的两个元素的下标，比如从数组[1,3,5,7,8]中找出和为 8 的两个元素的下标分别为(0,3)和(1,2)。
	for i := 0; i < len(aa); i++ {
		for j := i + 1; j < len(aa); j++ {
			if aa[i]+aa[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}
