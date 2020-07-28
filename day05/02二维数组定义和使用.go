package main

import "fmt"

func main0201() {
	// 二维数组的定义 var 数组名 [行个数][列个数] 数据类型
	var arr [3][4]int

	// 通过行和列下标找到具体元素进行赋值
	arr[1][2] = 3
	arr[2][1] = 4
	fmt.Println(arr)

	var arr1 [3][4]int = [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 2, 4, 5}}
	fmt.Println(arr1)

	arr2 := [3][4]int{{1, 2}, {3, 2: 9}, {4, 2: 9}}
	fmt.Println(arr2)

	for i := 0; i < len(arr2); i++ {
		// len(arr2[i]) 由于二维数组由行和列组成第一次循环得到的结果仅为每一行的内容被包裹在一个大数组中，所以此处 arr2[i] 表示每一行数组
		for j := 0; j < len(arr2[i]); j++ {
			// arr2[i][j] 由于二维数组由行和列组成，该语法表示取行与列交叉点的元素值
			fmt.Print(arr2[i][j], "  ")
		}
		fmt.Println()
	}
}
