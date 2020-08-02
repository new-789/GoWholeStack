package main

import "fmt"

type Stu struct {
	id   int
	name string
	age  int
}

func main1003() {
	// 定义结构体切片
	var stu []Stu = make([]Stu, 3)
	// 定义结构体切片指针，并进行赋值
	var p *[]Stu = &stu

	// 通过结构体切片指针变量往切片中的结构体类型元素添加数据
	(*p)[0] = Stu{101, "小猪佩奇", 10}
	(*p)[1] = Stu{102, "野猪佩奇", 15}
	(*p)[2] = Stu{103, "猪刚鬣", 15}

	// 通过普通 for 循环打印结构体切片指针中的每一个元素
	for i := 0; i < len(*p); i++ {
		fmt.Println("具体的值：", (*p)[i])
	}
	fmt.Println("------------ 分割线 ------------")
	// 通过普通 for range 循环打印结构体切片指针中的每一个元素
	for i, v := range *p {
		fmt.Println("下标：", i, " 具体的值：", v)
	}
}

// 结构体切片指针操作切片中的结构体
func main1002() {
	var stu []Stu = make([]Stu, 3)
	var p *[]Stu = &stu // 结构体切片指针
	//*p = append(*p, Stu{101, "小猪佩奇", 10})
	//fmt.Println(stu)

	// 通过结构体切片指针变量往切片中的结构体类型元素添加数据
	(*p)[0] = Stu{101, "小猪佩奇", 10}
	(*p)[1] = Stu{102, "野猪佩奇", 15}
	(*p)[2] = Stu{103, "猪刚鬣", 15}
	fmt.Println("添加内容后结果: ", stu)
}

// 结构体切片指针的定义
func main1001() {
	// 定义一个结构体切片并初始化,必须指定一个初始容量
	var stu []Stu = make([]Stu, 3)
	// 定义一个结构体切片指针，并与结构体名建立关系，即将结构体名指向的内存地址赋值给指针变量
	var p *[]Stu = &stu

	fmt.Printf("查看指针变量p的类型：%T\n", p)
}

///
