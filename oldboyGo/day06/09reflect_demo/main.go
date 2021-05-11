package main

import (
	"fmt"
	"reflect"
)

// reflect

type cat struct {
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
	fmt.Printf("type Name: %v type kind:%v\n", v.Name(), v.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	var k = v.Kind() // 获取具体值的种类
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int value:%d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32 value:%f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64 value :%f\n", float64(v.Float()))
	}
}

// 通过反射
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) // 修改的是副本，reflect 包会引发 panic
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// v.Elem() 获取指针指定对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	var a float32 = 3.14
	reflectType(a)
	var b int64 = 100
	reflectType(b)

	var c = cat{}
	reflectType(c)

	// valueOf
	reflectValue(a)

	// 设置值
	reflectSetValue1(&b)
	fmt.Println(b)
	reflectSetValue2(&b)
	fmt.Println(b)
}
