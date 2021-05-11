package main

import "fmt"

// 结构体是值类型
type person struct {
	name, gender string
}

// Go 语言中函数传参数永远传的是拷贝
func f(x person) {
	x.gender = "男" // 此处修改的是副本的 gender
}

func f1(x *person) {
	// 语法糖：自动根据指针找对应的变量,完整写法：(*x).gender = "妖"
	x.gender = "妖" // 根据内存地址找到原变量修改就是原变量的值
}

func main() {
	var p person
	p.name = "琳琳"
	p.gender = "女"
	f(p)
	fmt.Println(p.gender) // 女

	f1(&p)
	fmt.Println(p.gender) // 妖
	// 结构体指针1
	var p2 = new(person)
	(*p2).name = "沙弥"
	p2.gender = "男"
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", p2)  // p2 的值保存的是一个内存地址，即 person 结构体的内存地址
	fmt.Printf("%p\n", &p2) // 求 p2 的内存地址

	// 结构体指针2:
	// 2.1: key value 初始化
	var p3 = &person{
		name: "悟能",
		// gender: "男",
	}
	fmt.Printf("%#v\n", p3)

	// 使用值列表的形式初始化，值的顺序要和结构体定义时字段的顺序一致
	p4 := &person{
		"白龙",
		"马",
	}
	fmt.Printf("%#v\n", p4)
}
