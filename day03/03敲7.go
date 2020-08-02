package main

import "fmt"

func main0301() {
	for i := 1; i <= 100; i++ {
		if i%10 == 7 || i/10 == 7 || i%7 == 0 {
			fmt.Printf("%d 敲桌子\n", i)
		} else {
			fmt.Printf("%d 过\n", i)
		}
	}
}
