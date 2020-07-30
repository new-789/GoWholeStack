package main

import "fmt"

type Stud struct {
	id    int
	name  string
	score [3]int // 处理将数组作为成员的类型
}

func main0801() {
	// 定义一个结构体切片，用来将无名学生信息保存在该切片中
	student := []Stud{
		Stud{101, "小米", [3]int{31, 32, 33}},
		Stud{102, "小明", [3]int{41, 42, 43}},
		Stud{103, "小红", [3]int{51, 52, 53}},
		Stud{104, "小花", [3]int{61, 62, 63}},
		Stud{105, "小王", [3]int{71, 72, 73}},
	}

	// 通过双层循环求出每个学生的成绩信息， 方式一
	for _, v := range student { // 第一层循环遍历出结构体切片中的所有结构体信息
		sumScore := 0 // 用来保存总成绩
		avgScore := 0 // 用来保存平均成绩
		// 第二层循环遍历出结构体中 score 成绩为数组中的每一门成绩信息
		for _, data := range v.score {
			sumScore += data
			avgScore = sumScore / len(v.score)
		}
		fmt.Printf("%s 的总成绩为：%d, 他的平均成绩为：%d\n", v.name, sumScore, avgScore)
	}

	fmt.Println("=====================")

	// 方式二:
	for i := 0; i < len(student); i++ { // 第一层循环遍历出结构体切片中的所有结构体信息
		sumScore := 0 // 用来保存总成绩
		avgScore := 0 // 用来保存平均成绩
		// 第二层循环遍历出结构体中 score 成绩为数组中的每一门成绩信息
		for j := 0; j < len(student[i].score); j++ {
			sumScore += student[i].score[j]
			avgScore = sumScore / len(student[i].score)
		}
		fmt.Printf("%s 的总成绩为：%d, 他的平均成绩为：%d\n", student[i].name, sumScore, avgScore)
	}
}
