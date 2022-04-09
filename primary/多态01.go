package main

import "fmt"

type person1 struct {
	name string
	sex string
	age int
}
type student1 struct {
	person1
	score int
}

type teacher1 struct {
	person1
	subject string
}

func (stu * student1)SayHello(){
	fmt.Printf("我是%s 今年%d岁, 成绩%d\n", stu.name, stu.age, stu.score)
}
func (t * teacher1)SayHello(){
	fmt.Printf("我是%s 今年%d岁, 教授%s\n", t.name, t.age, t.subject)
}

// 接口实现, 接口属于全局

type Personer interface {
	SayHello()
}
// 多态实现 将接口类型作为函数参数实现
func Sayhello(p Personer){
	p.SayHello()
}

func main(){
	var p Personer

	p = &student1{person1{"小明", "男", 18}, 78}
	Sayhello(p)
	p = &teacher1{person1{"大师", "男", 43},"理论"}
	Sayhello(p)
}
