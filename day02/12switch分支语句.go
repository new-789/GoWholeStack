package main

import "fmt"

func main1201() {
	var lv string
	var wages int = 5000
	fmt.Printf("请输入评级>>:")
	fmt.Scan(&lv)

	if lv == "A" {
		wages += 500
	} else if lv == "B" {
		wages += 200
	} else if lv == "C" {
		wages = 5000
	} else if lv == "D" {
		wages -= 200
	} else {
		wages -= 500
	}
	fmt.Printf("李四来年的工资为:%d\n", wages)
}

func main1202() {
	var wages int = 5000
	var lv string
	fmt.Printf("请输入评级>>:")
	fmt.Scan(&lv)
	switch lv {
	case "A":
		wages += 500
	case "B":
		wages += 200
	case "D":
		wages -= 200
	case "E":
		wages -= 500
	case "C":
		wages = 5000
	default:
		fmt.Println("输入错误")
	}
	fmt.Printf("李四来年的工资为:%d\n", wages)
}

func main1203() {
	var w int
	fmt.Scan(&w)
	switch w {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6:
		fmt.Println("星期六")
	case 7:
		fmt.Println("星期天")
	default:
		fmt.Println("输入错误")
	}
}

func main() {
	var score int
	fmt.Printf("输入分数:")
	fmt.Scan(&score)

	switch score / 10 { // 此处支持一个 int 类型的表达式语句
	case 10:
		// 让 switch 执行下一个分支中的代码，即如果满足 10 和 9 之后执行的代码相同则直接执行 9 分支中的代码即可，如果不写则执行到下一个分支就会自动停止
		fallthrough
	case 9:
		fmt.Println("A")
	case 8:
		fmt.Println("B")
	case 7:
		fmt.Println("C")
	case 6:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}
}
