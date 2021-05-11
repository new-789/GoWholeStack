package main

import "fmt"

// 方法

// 标识符：变量名、函数名、类型名、方法名
// Go 语言中如果标识符首字母是大写的，则表示对外部包可见(公有的)
// 并且 Go 语言中如果首字母是大写的标识符，则必须对其进行添加注释，
// 注释的格式为： // 标识符(空格)注释的内容

// Dog 这是一个 Dog 结构体
type dog struct {
	name string
}

type person struct {
	name string
	age  int
}

// 构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

// 方法的定义：方法是作用域特定类型的一个函数
// 接收者：表示的是调用该方法的具体类型，多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s: 汪汪汪~\n", d.name)
}

// 使用值接收者：传的是拷贝内容
func (p person) guonian() {
	p.age++
}

// 指针接收者：传的是内存地址
func (p *person) zhenguonian() {
	p.age++
}

func (p *person) dreams() {
	fmt.Println("学会 Go 日薪越亿")
}

func main() {
	d1 := newDog("旺财")
	d1.wang()

	p1 := newPerson("悟能", 2000)
	// p1.wang() 不能调用 dog 类型接收者的方法
	fmt.Println(p1.age) // 2000
	p1.guonian()
	fmt.Println(p1.age) // 2000 值接收者不能更改结构体中字段的值
	p1.zhenguonian()
	fmt.Println(p1.age) // 2001
	p1.dreams()
}
