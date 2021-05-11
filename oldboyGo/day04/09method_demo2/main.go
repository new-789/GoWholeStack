package main

import "fmt"

// 给自定义类型添加方法
// 不能给别的包里面的类型添加方法，只能给自己包里的类型添加方法

// MyInt 给 int 自定义一个类型别名
type MyInt int

func (m MyInt) hello() {
	fmt.Println("我是一个 int")
}

type person struct {
	name string
	age  int
}

func main() {
	// 声明一个 int32 类型的变量，它的值是10
	// 方法一：
	// var x int32
	// x = 10
	// 方法2
	// var x int32 = 10
	// fmt.Println(x)
	// 方法三：
	// var x = int32(10)
	// 方法四：
	// x := int32(10)

	// 声明一个 MyInt 类型的变量 m 它的值是100
	// 方法一
	// var m MyInt
	// m = 100
	// 方法二
	// var m MyInt = 100
	// 方法三
	// var m = MyInt(100)
	// 方法四
	// m := MyInt(100) // 强制类型转换
	// fmt.Println(m)

	// 问题2：结构体初始化
	// 方法一
	var p person // 声明了一个 person 类型的变量
	p.name = "悟净"
	p.age = 18
	fmt.Println(p)
	var p1 person
	p1.name = "悟能"
	p1.age = 28
	fmt.Println(p1)
	// 方法二：
	// 键值对
	var p2 = person{
		name: "悟空",
		age: 38，
	}
	fmt.Println(p2)
	// 值列表
	var p3 = person{
		"师傅",
		88,
	}
	fmt.Println(p3)

	// var m = MyInt(100) // 初始化自定义类型
	// m.hello()
}

// 问题三：为什么要有构造函数，构造函数仅仅是用来初始化结构体使用，初始化后返回一个结构体变量
func nrePerson(name string, age int) *person {
	// 别人调用我，能够给调用人一个 person 类型的变量
	return &person{
		name: name, 
		age: age,
	}
}