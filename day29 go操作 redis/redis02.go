package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main2902() {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("链接数据库错误", err)
		return
	}
	defer conn.Close()
	/*
		// Do 实行 set 设置元素操作
		rel, err := conn.Do("set", "a1", "hello world")
		if err != nil {
			fmt.Println("Do 执行 set 操作错误", err)
			return
		}
		// string 方法使用
		rel, err = redis.String(conn.Do("get", "a1"))
		if err != nil {
			fmt.Println("Do 执行 get 操作错误", err)
			return
		}

		// 通过反射的方式获取需要的结果
		// rel, err = conn.Do("get", "a1")
		// fmt.Println(string(rel.(string)))
	*/

	// Do 执行 get 获取元素操作
	// rel, err = redis.Strings(conn.Do("mget", "name", "age", "class"))
	// if err != nil {
	// 	fmt.Println("Do 执行 get 操作错误", err)
	// 	return
	// }

	// values 方法使用
	rel, err := redis.Values(conn.Do("mget", "name", "age", "class"))
	// Scan 转换字符串数组并获取指定类型的值
	var name string
	var age int
	var class string
	redis.Scan(rel, &name, &age, &class)

	fmt.Println(name, age, class)
}
