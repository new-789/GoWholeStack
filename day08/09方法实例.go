package main

import "fmt"

// 方法一，打招呼
func (s *Student1) SayHello() {
	fmt.Printf("大叫好，我的名字叫: %s, 今年: %d 岁, 我是: %s 同学\n", s.name, s.age, s.sex)
}

// 方法二：打印成绩
func (s *Student1) PrintScore() {
	sum := s.cscore + s.mscore + s.escore
	fmt.Printf("我叫: %s, 本次考试的总成绩为:%d, 平均得分为：%d\n", s.name, sum, sum/3)
}
