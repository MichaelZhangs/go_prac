package main

import "fmt"

type person struct {
	id   int
	name string
	age  int
	sex  string
}
type student struct {
	//p     person   // 结构体变量 实名
	person
	score int
}
func (p *person) SayHello(){

	fmt.Println("我是"+p.name+ "年龄", p.age)
}
func (s *student) SayHello2(){
	fmt.Println("2 我是"+s.name+ "年龄", s.age)
}

func main(){
	var stu student = student{person{102,"小明", 19,"男"}, 92}
	//实名下的初始化
	//stu.p.age = 18
	//stu.p.id = 101
	//stu.p.sex = "女"
	//stu.p.name = "小红"
	//stu.score = 88
	stu.SayHello2()
	//stu.p.SayHello()

	fmt.Println(stu)
}
