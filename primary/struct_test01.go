package main

import (
	"fmt"
	"unsafe"
)

type Course struct {
	Name string
	Price int
	Url string
}

func (c Course)printInfoCourse(){
	c.Price = 400
	fmt.Println(c.Price, c.Name, c.Url)
}

func test(c *Course)  {
	c.Name = "scrapy"
}

func main() {

	var c Course = Course{"django", 100, "https://www.baidu.com"}
	fmt.Println(c.Name, c.Price, c.Url)
	// 结构体定义的名称，首字母大写则 全局可以访问，小写则对引用该包的其他文件不可访问
	fmt.Println("size", unsafe.Sizeof(c))
	c2 := &Course{"django", 100, "https://www.baidu.com"}
	fmt.Println(c2.Name, c2.Url)
	fmt.Println("size",unsafe.Sizeof(c2))
	test(c2)
	fmt.Println(c2.Name)
	c3 := &Course{}
	fmt.Println("size",unsafe.Sizeof(c3))
	fmt.Println(c3.Price)
	c4 := new(Course)
	fmt.Println(c4.Price)
	fmt.Println("size",unsafe.Sizeof(c4))
	c2.printInfoCourse()
	Course.printInfoCourse(*c2)
}