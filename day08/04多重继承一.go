package main

import "fmt"

type testA struct {
	name string
	id   int
}

type testB struct {
	testA //
	sex   string
	age   int
}

type testC struct {
	testB
	score int
}

func main0401() {
	var t testC = testC{testB{testA{"麻子", 201}, "男", 98}, 99}
	//t.testB.testA.name = "麻子"
	//t.name = "王二"
	//t.id = 201
	//t.sex = "男"
	//t.age = 30
	//t.score = 100

	fmt.Println(t)
}
