package main

import (
	"encoding/json"
	"fmt"
)

// 复习
type person struct {
	name string
	age  int
}

// func sum(x, y int) (ret int) {
// 	return x + y
// }

// 构造函数
func newPerson(n string, i int) person {
	return person{
		name: "悟空",
		age:  1000,
	}
}

// 方法：
// 接收者使用对用类型的首字母小写
// 指定了接收者之后，只有接收者这个类型的变量调用这个方法
func (p *person) dream(str string) {
	fmt.Printf("%s 的梦想是%s\n", p.name, str)
}

// func (p person) guonian() {
// 	p.age++ // 此处是 p1 的副本，所以改的是副本
// }

// 指针接收者
// 1、需要修改值时使用指针接收者
// 2、结构体本身比较大时拷贝的内存开销比较大时也要使用指针接收者
// 3、保持一致性：如果有一个方法使用了指针接收者，其他的方法为了统一也要使用指针接收者
func (p *person) guonian() {
	p.age++ //
}

func main() {
	p1 := person{"悟净", 777}
	p1.dream("挑好担")
	p2 := person{"悟能", 888}
	p2.dream("每天都能看见美女")

	fmt.Println(p1.age)
	p1.guonian()
	fmt.Println(p1.age)

	// 结构体嵌套
	type addr struct {
		province string
		city     string
	}

	type student struct {
		name string
		addr // 匿名嵌套别的结构体，就使用类型名做名称
	}

	// json 序列化
	type point struct {
		X int `json:"x"`
		Y int `json:"y"`
	}

	po1 := point{100, 200}
	b, err := json.Marshal(po1)
	if err != nil {
		fmt.Printf("marshal failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b))

	// json 反序列化：由字符串-->go语言中的结构体变量
	str := `{"x":10,"y":20}`
	var po2 point // 造一个结构体变量准备接收反序列化的值
	// 注意：第二个参数用来接收反序列化后的结构体变量名必须为指针类型
	err = json.Unmarshal([]byte(str), &po2)
	if err != nil {
		fmt.Printf("Unmarshal fialed, err:%v\n", err)
		return
	}
	fmt.Println(po2)
}
