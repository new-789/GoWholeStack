package main

import "fmt"

type JieKou interface {
	JiSuan() int
}

type FuLei struct {
	num1 int
	num2 int
}

type JiaFaZiLei struct {
	FuLei
}

type JianFaZiLei struct {
	FuLei
}

type GongChang struct{}

func (jiafa *JiaFaZiLei) JiSuan() int {
	return jiafa.num1 + jiafa.num2
}

func (jianfa *JianFaZiLei) JiSuan() int {
	return jianfa.num1 - jianfa.num2
}

func (g *GongChang) JiSuanJieGuo(num1, num2 int, yunsuanfu string) (value int) {
	var jiekou JieKou // 创建接口类型变量
	switch yunsuanfu {
	case "+":
		jiekou = &JiaFaZiLei{FuLei{num1, num2}}
	case "-":
		jiekou = &JianFaZiLei{FuLei{num1, num2}}
	default:
                // 此处代码有错误提示，待解决
		fmt.Println("sorry, 没有匹配到对应的方法，请检查您传入的参数是否正确")
	}
	value = DuoTai(jiekou)
	return
}

func DuoTai(jk JieKou) (value int) {
	value = jk.JiSuan()
	return
}

func main() {
	var gc GongChang
	value := gc.JiSuanJieGuo(100, 20, "-")
	fmt.Println(value)
}
