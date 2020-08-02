package main

import "fmt"

// 水仙花练习题
func main0201() {
	var d int
	for i := 100; i < 1000; i++ {
		a := i / 100     // 取百位上的数
		b := i / 10 % 10 // 取十位上的数，方式二：b := i % 100 /10
		c := i % 10      // 取各位上的数
		d = a*a*a + b*b*b + c*c*c
		if d == i {
			fmt.Printf("%d 是水仙花数\n", d)
		}
	}
}
