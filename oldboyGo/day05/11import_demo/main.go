package main

import (
	// github.com/GoWholeStack/oldboyGo/day05/10calc
	calc "github.com/GoWholeStack/oldboyGo/day05/10calc"
	"fmt"
)

func init() {
	fmt.Println("自动执行")
}

func main() {
	sum := calc.Add(10, 20)
	fmt.Println(sum)
}