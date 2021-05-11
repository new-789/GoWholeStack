package main

import (
	"encoding/json"
	"fmt"
)

// 反射

type student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"三藏", "age": 9999}`
	var s student
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s.Name, s.Age)
}
