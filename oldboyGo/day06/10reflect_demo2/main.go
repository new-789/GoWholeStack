package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name" sanzang:"哈哈"`
	Score int    `json:"score" sanzang:"啦啦"`
}

func main() {
	stu1 := student{
		Name:  "三藏",
		Score: 909,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过 for 循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		filed := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", filed.Name, filed.Index, filed.Type, filed.Tag.Get("sanzang"))
	}

	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("sanzang"))
	}
}
