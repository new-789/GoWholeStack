package main

import "fmt"

// 定义一个接口，并指定一个 Operate 方法，且带一个 int 类型的返回值
type Opter interface {
	Operate() int
}

// 父类
type Opt struct {
	num1 int
	num2 int
}

// 加法子类
type AddOpt struct {
	Opt
}

// 减法子类
type SubOpt struct {
	Opt
}

// 加法方法的实现
func (add *AddOpt) Operate() int {
	return add.num1 + add.num2
}

// 减法方法的实现
func (sub *SubOpt) Operate() int {
	return sub.num1 - sub.num2
}

func main0101() {
	// 基于继承方法和接口实现加减法运算
	var o Opter
	// 通过接口调用加法方法
	o = &AddOpt{Opt{10, 20}}
	value := o.Operate()
	fmt.Println(value)

	// 通过结构调用减法方法
	o = &SubOpt{Opt{100, 30}}
	result := o.Operate()
	fmt.Println(result)
}
