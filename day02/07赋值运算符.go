package main

import "fmt"

func main0701() {
	a := 10
	a += 5 // == a = a + 5
	fmt.Println(a)
	a -= 5 // == a = a - 5
	fmt.Println(a)
	a *= 5 // == a = a * 5
	fmt.Println(a)
	a /= 5 // == a = a / 5
	fmt.Println(a)
	a %= 5 // == a = a % 5
	fmt.Println(a)
}
