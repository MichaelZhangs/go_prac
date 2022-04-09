package main

import "fmt"

type person struct {
	name string
	age  int
	sex  string
}

type Student struct {
	// 通过匿名字段实现继承操作
	//person  // 结构体名称作为成员
	*person // 指针名
	id    int
	score int
}

func main() {
	//var stu Student = Student{person{"张三",19,"男"},101,78}
	//stu.id = 101
	//stu.name = "张三"
	////stu.person.name  = "李四"
	//stu.score = 100
	//stu.sex = "男"
	//stu.age = 18
	var stu Student = Student{&person{"张三",21,"女"},102,78}
	//stu.person = new(person)  // 必须先new
	//stu.name = "李四"
	//stu.score = 100
	//stu.sex = "男"
	//stu.age = 18
	//stu.id = 1
	fmt.Println(stu)
	fmt.Println(stu.name)
	fmt.Println(stu.age)
	fmt.Println(stu.sex)
}
