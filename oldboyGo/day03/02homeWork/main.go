package main

import "fmt"

// 作业
func main() {
	// 1. 判断字符串中汉字的数量
	// 难点：如何判断一个字符是汉字
	// s1 := "hello沙河小王子"
	// // 1. 一次拿到字符串中的字符
	// var count int
	// for _, v := range s1 {
	// 	// 2. 判断当前这个字是否为汉字
	// 	if unicode.Is(unicode.Han, v) {
	// 		count++
	// 	}
	// }
	// // 3. 将汉字出现的次数累加得到最终的结果
	// fmt.Println(count)

	// 2. how do you do 判断单词出现的次数
	// s2 := "how do you do"
	// ss := strings.Split(s2, " ")
	// m := make(map[string]int, 5)
	// for _, v := range ss {
	// 	_, ok := m[v]
	// 	if !ok {
	// 		m[v] = 1
	// 	} else {
	// 		m[v]++
	// 	}
	// }
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// 回文判断：字符串从左往右读从右往左读是一样的，那么就是回文，
	// 如上海自来水来自上海，山西运煤车煤运西山，黄山落叶松叶落山黄
	// 解题思路：
	// 将字符串中的字符拿出来放到一个 []rune 类型的切片中
	s1 := "山西运煤车煤运西山"
	r := make([]rune, 0, len(s1))
	for _, v := range s1 {
		r = append(r, v)
	}
	fmt.Println("[]rune:", r)

	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}
