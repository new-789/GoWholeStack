package main

import "fmt"

func main0601() {
	// 创建 map
	/*
		var m1 map[int]string  // 声明 map ，没有空间，只能直接存储 key == value

		m2 := map[int]string{}
		m2[101] = "hello"
		m2[102] = "world"
		fmt.Println(len(m2))
		fmt.Println("m2 = ", m2)

		m3 := make(map[int]string)
		m3[101] = "come"
		m3[102] = "brother"
		fmt.Println("m3 = ", m3)

		m4 := make(map[int]string, 5)  // 参数二为指定的长度，但可以自动扩容
		m4[188] = "china"
		m4[888] = "I Love"
		fmt.Println("m4 = ", m4)
	*/

	// 初始化 map
	/*
		var m5 map[int]string = map[int]string{109:"你好", 888:"路飞", 108: "日薪越亿"}
		fmt.Println("m5 = ", m5)

		m6 := map[int]string{800:"这是", 102:"无序", 3: "map"}
		fmt.Println("m6 = ", m6)
	*/

	// map 赋值操作
	/*
		m7 := make(map[int]string, 1)
		m7[100] = "天成"
		m7[20] = "酒店"
		m7[3] = "美女多"
		fmt.Println("m7 = ", m7)

		m7[3] = "去把妹"  // 此句操作会将原 map 中 key 为 3 的 map 元素进行替换
		m7[8] = "欢乐谷"
		fmt.Println("m7 = ", m7)
	*/

	// map 的遍历
	/*
		var m8 map[int]string = map[int]string{109: "你好", 888: "路飞", 108: "日薪越亿"}
		for k, v := range m8 {
			fmt.Printf("key：%d --> value：%s\n", k, v)
		}
		// range 循环字典时返回的 key / value ，默认可省略 value 进行输出
		for k := range m8 {
			fmt.Println("key: ", k)
		}
		// 忽略 key 进行遍历
		for _, v := range m8 {
			fmt.Printf("value:%s\n", v)
		}
	*/

	// 判断 map 中的 key 是否存在
	var m9 map[int]string = map[int]string{109: "你好", 888: "路飞", 108: "日薪越亿"}
	if v, has := m9[888]; !has { // m9[下标] 返回两个值，第一个为 value ，第二个为 bool 代表是否存在
		fmt.Printf("key if not find\n")
	} else {
		fmt.Println("value =", v, "has = ", has)
	}
}
