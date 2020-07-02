package main

import "fmt"

func main1001() {
	a := 10
	b := 20
	c := 30
	d := a + b*c
	e := (a + b) * c
	fmt.Println(d)
	fmt.Println(e)
}

func main1002() {
	a := 10
	b := 20
	c := 30

	fmt.Println(a+b >= c && !(b > c))
}
