package main

import "fmt"

type person struct {
	name string
	sex string
	age int
}
type student struct {
	person
	score int
}

type teacher struct {
	person
	subject string
}

func (stu * student)Sayhello(){
	fmt.Printf("我是%s 今年%d岁, 成绩%d\n", stu.name, stu.age, stu.score)
}
func (t * teacher)Sayhello(){
	fmt.Printf("我是%s 今年%d岁, 教授%s\n", t.name, t.age, t.subject)
}
//接口定义
// 接口定义了规则， 方法实现了规则
type Humaner interface {
	// 定义方法
	Sayhello()
}

func main(){
	var stu student = student{person{"小明","男",18},82}
	var t teacher = teacher{person{"大师","男",43},"武艺"}
	//stu.Sayhello()
	//t.Sayhello()

	var h Humaner
	h = &stu
	h.Sayhello()
	h = &t
	h.Sayhello()

}
