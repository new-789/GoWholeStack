package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// gob 包使用示例
//定义一个 Person 结构

type Person struct {
	Name string
	Age int
}

func main() {
	var minming Person
	minming.Name = "小明"
	minming.Age = 28
	var buffer bytes.Buffer // 编码后的数据放到 buf 中
	//使用 gob 进行序列化(编码)得到字节流
	// 1. 定义一个编码器
	encoder := gob.NewEncoder(&buffer)
	// 2. 使用编码器进行编码
	err := encoder.Encode(&minming)
	if err != nil {
		log.Panic("编码出错，小明不知去向")
	}
	fmt.Printf("编码后的数据：%v\n", buffer.Bytes())

	//使用 gob 进行反序列化
	// 1. 定义一个解码器
	// func NewDecoder(r io.Reader) *Decoder
	decoder := gob.NewDecoder(bytes.NewBuffer(buffer.Bytes()))
	// 2. 使用解码器对数据进行解码
	var daming Person // 解码后的数据存放到 daming 变量中
	err = decoder.Decode(&daming)
	if err != nil {
		log.Panic("解码出错")
	}
	fmt.Printf("解码后的数据：%v\n", daming)
}
