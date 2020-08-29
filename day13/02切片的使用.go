package main

import "fmt"

func main0201() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s := arr[1:5:7]

	fmt.Println("s = ", s)
	fmt.Println("len(s) = ", len(s))
	fmt.Println("cap(s) = ", cap(s))

	s2 := s[:6]
	fmt.Println("s2 = ", s2)
	fmt.Println("len(s2) = ", len(s2))
	fmt.Println("cap(s2) = ", cap(s2))
}

func main0202() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	s := arr[2:5:5] // 使用截取数组的方式初始化切片 s ，并指定了容量为 5
	fmt.Println("s = ", s)
	fmt.Println("len(s) = ", len(s))
	fmt.Println("cap(s) = ", cap(s))

	s2 := s[2:7] // 使用截取切片 s 的方式初始化切片 s2 ,注意此时截取的长度大于切片 s 的容量
	fmt.Println("s2 = ", s2)
	fmt.Println("len(s2) = ", len(s2))
	fmt.Println("cap(s2) = ", cap(s2))
}

// 创建切片
func main0203() {
	s1 := []int{1, 2, 3, 4}
	fmt.Println("s1 = ", s1)

	s2 := make([]int, 5, 10)
	fmt.Println("len(s2) =", len(s2), ", cap(s2) = ", cap(s2))

	s3 := make([]int, 7)
	fmt.Println("len(s3) =", len(s3), ", cap(s3) = ", cap(s3))
}

func main0204() {
	//s := []int{1, 2, 3, 4}  // 创建一个有初始值的切片
	//s = append(s,888)
	//s = append(s,888)
	//s = append(s,888)
	//s = append(s,888)
	//s = append(s,888)
	//fmt.Println(s)

	s := make([]int, 0, 1)
	c := cap(s)
	for i := 0; i < 100; i++ {
		s = append(s, i)
		for n := cap(s); n > c; {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}
}
