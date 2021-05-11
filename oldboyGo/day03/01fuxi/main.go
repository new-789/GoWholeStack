package main

import "fmt"

// day03复习
func main() {
	// var name string
	// name = "联想"
	// fmt.Println(name)
	// var ages [30]int // 声明了一个变量 ages，它是 [30]int 类型
	// fmt.Println(ages)
	// ages = [30]int{1, 2, 3, 4, 5, 6, 7, 8, 9} // 数组初始化方式一
	// fmt.Println(ages)
	// var ages2 = [...]int{1, 2, 3, 4, 5} // 数组初始化方式二
	// fmt.Println(ages2)
	// var ages3 = [...]int{1: 100, 99: 200} // 数组初始化方式三,指定索引初始化
	// fmt.Println(ages3)

	// // 二维数组
	// // var a1 [...][2]int
	// // fmt.Println(a1)
	// // 多维数组只有最外层可以使用 ...
	// var a1 = [...][2]int{
	// 	[2]int{1, 2},
	// 	[2]int{3, 4},
	// 	[2]int{5, 6},
	// }
	// fmt.Println(a1)

	// // 数组是值类型
	// x := [3]int{1, 2, 3}
	// y := x         // 将 x 的值拷贝一份给了 y
	// y[1] = 200     // 修改了副本 y ，并不影响 x 数组中的值
	// fmt.Println(x) // [1,2,3]
	// modify(x)
	// fmt.Println(x) // [1,2,3]

	// ##########################################################
	// 切片(sile)：
	// var s1 []int // 没有分配内存，== nil
	// fmt.Println(s1)
	// fmt.Println(s1 == nil) // true
	// s1 = []int{1, 2, 3}    // 切片初始化方式一
	// fmt.Println(s1)
	// // make 初始化切片，分配内存
	// s2 := make([]bool, 2, 4)
	// fmt.Println(s2) // [false false]
	// s3 := make([]int, 0, 4)
	// fmt.Println(s3 == nil) // false

	// s4 := []int{1, 2, 3} // [1 2 3]
	// s5 := s4
	// var s6 = make([]int, 3, 3)
	// copy(s6, s4)    // copy 的前提是两个切片的长度一致，才能完成全部拷贝，否则只会拷贝新切片现有长度对应的元素
	// fmt.Println(s5) // [1 2 3]
	// s5[1] = 200
	// fmt.Println(s5) // [1 200 3]
	// fmt.Println(s4) // [1 200 3]
	// fmt.Println(s6) // [1 2 3]

	// var s1 []int // nil
	// // s1 = make([]int, 1, 2)
	// // s1[0] = 100
	// s1 = append(s1, 100) // append 会自动初始化切片
	// fmt.Println(s1)

	// 指针
	// Go 语言中的指针只能读不能修改，不能修改指针变量指向的地址
	// addr := "娜娜"
	// addrP := &addr
	// fmt.Println(addrP) // 内存地址
	// fmt.Printf("%T\n", addrP)
	// addrV := *addrP // 根据内存地址取值
	// fmt.Println(addrV)

	// map
	var m1 map[string]int
	fmt.Println(m1 == nil)
	m1 = make(map[string]int, 10)
	m1["娜娜"] = 200
	fmt.Println(m1)
	fmt.Println(m1["积极"]) // 如果 key 不存在，返回的是对应类型的零值
	// 如果返回值是布尔型，通常使用 ok 去接收
	v, ok := m1["娜娜"]
	if !ok {
		fmt.Println("没有娜娜这个人")
	} else {
		fmt.Println("娜娜的分数是：", v)
	}
	delete(m1, "娜娜") // 删除 map 中指定 key 值，如果不存在则什么都不做
	fmt.Println(m1)
	fmt.Println(m1 == nil) // 已经开辟了内存空间，所以不等于 nil，只有刚创建没有初始化前 map 才等于 nil
}

func modify(a [3]int) {
	// go 语言中的函数的参数，传递的值都是（ctrl+c, ctrl+v）
	a[1] = 100 // 此处修改的是副本的值
}
