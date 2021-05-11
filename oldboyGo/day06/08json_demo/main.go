package main

import (
	"encoding/json"
	"fmt"
)

// json
type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"三藏", "age":9000}`
	var p person
	json.Unmarshal([]byte(str), &p)
	fmt.Println(p.Name)
	fmt.Println(p.Age)
}
