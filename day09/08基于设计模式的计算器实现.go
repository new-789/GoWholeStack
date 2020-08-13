package main

import "fmt"

// 设计模式的实现，
type OptFactory struct{}

// 工厂模式实现加减法，接收两个int类型参数(num1,num2)，和一个运算符参数(op),并返回一个 int 类型数据
func (factory *OptFactory) OptCalc(num1, num2 int, op string) (value int) {
	var o Opter1
	// 通过对运算符进行断言实现分类运算
	switch op {
	case "+":
		// 将 AddOpt1 加法结构体对象的地址赋值给接口变量
		o = &AddOpt1{Opt1{num1, num2}}
	case "-":
		// 将 SubOpt1 减法结构体对象的地址赋值给接口变量
		o = &SubOpt1{Opt1{num1, num2}}
	}
	// 通过调用 NewOpt 多态函数实现接口的操作
	value = Calculation(o)
	return
}

// 多态实现
func Calculation(o Opter1) (value int) {
	// 通过船体过来的接口变量统一处理，调用不同结构体(如加、减、乘、除结构体)绑定的方法
	value = o.Operate1()
	return
}

// 定义一个接口，并指定一个 Operate 方法，且带一个 int 类型的返回值
type Opter1 interface {
	Operate1() int
}

// 父类
type Opt1 struct {
	num1 int
	num2 int
}

// 加法子类
type AddOpt1 struct {
	Opt1
}

// 减法子类
type SubOpt1 struct {
	Opt1
}

// 加法方法的实现
func (add *AddOpt1) Operate1() int {
	return add.num1 + add.num2
}

// 减法方法的实现
func (sub *SubOpt1) Operate1() int {
	return sub.num1 - sub.num2
}

func main0801() {
	//基于继承、方法、接口、多态和设计模式
	factory := OptFactory{}
	// 根据传参的不同，即实现了不同的方法，如传递 + 号，则执行加法运算方法，传递的为 - 号则执行的是减法运算
	value := factory.OptCalc(20, 60, "-")
	fmt.Println("运算结果：", value)
}
