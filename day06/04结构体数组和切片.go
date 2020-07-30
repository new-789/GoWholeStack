package main

import "fmt"

type Stu struct {
	id    int
	name  string
	sex   string
	age   int
	score int
	addr  string
}

func main0401() {
	// 定义结构体数组
	var arr [3]Stu
	// 结构体数组赋值语法：结构体数组名[下标].成员 = 值

	fmt.Println("Please input corresponding value:")
	// 通过循环的方式给结构体数组中的成员进行赋值
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i].id, &arr[i].name, &arr[i].sex, &arr[i].age, &arr[i].score, &arr[i].addr)
	}

	// 打印输出结构体数组中结构体信息
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

// 结构体数组的应用
func main0402() {
	var arr [5]Stu

	fmt.Println("Please input corresponding value:")
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i].id, &arr[i].name, &arr[i].sex, &arr[i].age, &arr[i].score, &arr[i].addr)
	}

	// 对结构体数组通过冒泡排序进行排序
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			// 通过结构体成员进行比较，然后进行交换以达到排序的目的
			if arr[j].age > arr[j+1].age {
				// 结构体数组中的的结构体相互赋值操作
				arr[j], arr[j+1] = arr[j+1], arr[j]

				// 结构体成员一次交换,不推荐使用
				//arr[j].id, arr[j+1].id = arr[j+1].id, arr[j].id
				//arr[j].name, arr[j+1].name = arr[j+1].name, arr[j].name
				//arr[j].sex, arr[j+1].sex = arr[j+1].sex, arr[j].sex
				//arr[j].age, arr[j+1].age = arr[j+1].age, arr[j].age
				//arr[j].score, arr[j+1].score = arr[j+1].score, arr[j].score
				//arr[j].addr, arr[j+1].addr = arr[j+1].addr, arr[j].addr
			}
		}
	}
	fmt.Println("----------------------------------")
	// 打印输出结果
	for i := 0; i < len(arr); i++ {
		fmt.Println("排序后的结果: ", arr[i])
	}
}

// 自动类型推导式的结构体数组赋值方式
func main0403() {
	arr := [2]Stu{
		{101, "孙尚香", "女", 20, 90, "江东"},
		{102, "黄月英", "女", 25, 70, "襄阳"},
	}

	for i, v := range arr {
		fmt.Println(i, v)
	}
}

// 结构体切片
func main0404() {
	// 通过自动推导类型语法创建结构体切片变量
	arr := []Stu{
		{101, "孙尚香", "女", 20, 90, "江东"},
		{102, "黄月英", "女", 25, 70, "襄阳"},
	}

	// 在结构体切片中增加数据
	arr = append(arr, Stu{103, "黄月英", "女", 25, 70, "襄阳"})

	for i, v := range arr {
		fmt.Println(i, v)
	}
}
