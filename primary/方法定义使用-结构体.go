package main

import "fmt"

type stu struct {
	sex  string
	id   int
	name string
}

func (s stu)PrintInfo() stu {
	s.name = "小明"

	return s
}

func main() {
	var s stu = stu{"女", 101, "小红"}
	fmt.Println(s.PrintInfo())
}
