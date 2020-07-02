package main

import "fmt"

func main1101() {
	var score int
	fmt.Scan(&score)
	if score > 720 {
		fmt.Println("考入了清华")
	} else {
		fmt.Println("回去再读几年")
	}
}
func main1102() {
	var score int
	fmt.Printf("请输入分数>>:")
	fmt.Scan(&score)

	if score > 700 {
		fmt.Println("清华我来了")
		if score > 720 {
			fmt.Println("我要学习挖掘机")
		} else if score > 710 {
			fmt.Println("我要学习美容美发")
		} else {
			fmt.Println("我要学习计算机")
		}
	} else if score > 680 {
		fmt.Println("北大我来了")
		if score > 690 {
			fmt.Println("我要学习卜卦")
		} else if score > 685 {
			fmt.Println("我要学习盗墓")
		} else {
			fmt.Println("我要学习种植")
		}
	} else if score > 650 {
		fmt.Println("乖乖到蓝翔去吧")
	} else {
		fmt.Println("回去继续努力")
	}
}

func main1103() {
	var a int = 10
	if a > 5 {
		fmt.Println(a)
	}
	// 采用就近原则，找到上面尚未配对的 if 进行匹配操作
	if a > 8 {
		fmt.Println(a)
	} else {
		fmt.Println("Asdasd")
	}
}

func main1104() {
	// 求三只佩奇那个最重
	var a, b, c int
	fmt.Println("请输入三只佩奇的体重>>:")
	fmt.Scan(&a, &b, &c)

	if a > b {
		if a > c {
			fmt.Println("a重")
		} else {
			fmt.Println("c重")
		}
	} else {
		if b > c {
			fmt.Println("b重")
		} else {
			fmt.Println("c重")
		}
	}
}
