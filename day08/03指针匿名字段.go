package main

import "fmt"

type People struct {
	name string
	age  int
	sex  string
}
type Stu struct {
	*People // 指针匿名字段
	id      int
	score   int
}

func main0301() {
	var stu Stu = Stu{&People{"杨过", 25, "男"}, 101, 99}
	// 使用 new() 方法为指针匿名字段创建空间
	stu.People = new(People)
	// 初始化内嵌结构体中的字段
	stu.name = "周芷若"
	stu.sex = "尼姑"
	stu.age = 25
	// 初始化自身字段
	stu.id = 101
	stu.score = 100

	// 通过嵌套结构体类型变量名直接调用内嵌匿名字段结构体中的字段名可获取到初始化的值
	fmt.Println(stu.name)
	fmt.Println(stu.age)
	fmt.Println(stu.sex)
}
