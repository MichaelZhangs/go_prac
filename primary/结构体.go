package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  bool // true 为男
	age  int
	addr string
}

func main() {
	stu := Student{1, "张三", true, 18, "广东"}
	fmt.Println(stu)
	var s Student
	s.id = 2
	s.name = "李四"
	stu2 := Student{1, "张三", true, 18, "广东"}
	fmt.Println(stu == stu2)
	fmt.Printf("%p %p", &stu, &stu2)

}
