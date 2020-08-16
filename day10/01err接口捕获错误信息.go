package main

import (
	"errors"
	"fmt"
)

func Test(a, b int) (value int, err error) {
	// 0 不能作为除数
	if b == 0 {
		err = errors.New("除数不能为 0 ，请重新指定一个参数")
	} else {
		value = a / b
	}
	return
}

func main0101() {
	value, err := Test(10, 0)
	// err 如果不等于空表示有错误信息
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}
