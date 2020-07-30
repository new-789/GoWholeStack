package main

import "fmt"

func main0101() {
	var arr [20]byte
	fmt.Print("Please input str>>：")
	for i := 0; i < len(arr); i++ {
		// 循环获取用户输入的数据
		fmt.Scanf("%c", &arr[i])
	}

	// 通过循环往 map 中添加数据
	m := make(map[byte]int, 1)
	for i := 0; i < len(arr); i++ {
		// 往 map 中添加数据，如果遇到相同的则会将其原来的值覆盖，且值会自增即为出现的次数
		// m[arr[i]] 表示将 arr 数组中的每一个值当做 key 存入数组
		// ++ 则为该字符出现的次数，因为没有出现则默认值为0，出现一次就会自增一次
		m[arr[i]]++
	}
	// 打印输出 map 中的数据，查看每个字符出现的次数
	for k, v := range m {
		if v > 0 {
			fmt.Printf("%c  %d\n", k, v)
		}
	}
}
