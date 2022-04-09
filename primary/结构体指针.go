package main

import "fmt"

type Student struct {
	id   int
	name string
	age  int
	sex  string
}

func main(){
	// 定义结构体变量
	// var 结构体名 结构体数据类型
	var stu Student = Student{1,"张三", 18, "男"}
	fmt.Println(stu.name)

	// 结构体指针
	var p * Student
	p = &stu
	p.name = "小明"
	fmt.Printf("%p\n", p)
	fmt.Printf("%p\n", &stu)

	fmt.Println(p)
	//结构体切片
	var st []Student = make([]Student, 3)
	s := &st
	fmt.Printf("%T\n", s)
	//*s = append(*s, Student{101, "李四",21, "男"})
	(*s)[0] = Student{101, "李四",21, "男"}
	fmt.Println(st)

}
