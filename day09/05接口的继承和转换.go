package main

import "fmt"

type Humaner1 interface { // 子集：被继承的接口
	SayHi()
}

type Personer interface { // 超集：继承其它接口的接口
	Humaner1 // 继承的 Humaner1 接口,一组子集的集合
	Sing(string)
}

type student1 struct {
	name string
	age  int
	sex  string
}

func (s *student1) SayHi() {
	fmt.Printf("Hello，我是：%s, 性别：%s, 今年：%d 岁\n", s.name, s.sex, s.age)
}

func (s *student1) Sing(SongTitle string) {
	fmt.Printf("大家好，我叫：%s，下面我给大家唱一首：《%s》 \n", s.name, SongTitle)
}

func main0501() {
	var h Humaner1 // 子集
	h = &student1{"韩红", 40, "女"}
	//h.SayHi()

	var p Personer // 超集
	p = &student1{"张柏芝", 20, "女"}
	p.Sing("青春不悔")
	//p.SayHi()  // 通过超集变量调用子集接口中的方法

	// 将超集转换为子集
	h = p
	// p = h // err, 不允许将子集转换为超集
	h.SayHi()
}
