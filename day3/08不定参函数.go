package main

import "fmt"

func test(args ...int) {
	for i := 0; i < len(args); i++ {
		fmt.Println("下标", i, "值", args[i])
	}
}

func plus(ar ...int) {
	sum := 0
	//for i:=0; i < len(ar); i++ {
	//	sum += ar[i]
	//}
	for _, data := range ar {
		sum += data
	}
	fmt.Println(sum)
}
func main0801() {
	//test(1,2,3,4,)
	plus(1, 2, 3, 3)
	plus(1, 2, 3, 4, 5, 6, 7)
	plus(3, 5, 7, 8, 1, 9, 10)
}
