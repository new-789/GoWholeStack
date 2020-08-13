package main

import "fmt"

func main0601() {
	var i interface{} = 10
	fmt.Println(i)
	fmt.Printf("int 类型：%p\n", &i)
	i = "hello world"
	fmt.Println(i)
	fmt.Printf("string 类型：%p\n", &i)
	var arr [3]int = [3]int{1, 2, 3}
	i = arr
	fmt.Println(i)
	fmt.Printf("array 类型：%p\n", &i)
}

// 空接口也可以组合成一个复合类型数据
func main0602() {
	// 定义空接口切片
	var i []interface{}
	i = append(i, 1, 2, "hello", "哈哈", [3]int{1, 2, 3})

	//for j:=0;j<len(i) ;j++  {
	//	fmt.Println(i[j])
	//}

	for idx, v := range i {
		fmt.Println(idx, v)
	}
}
