package main

import "fmt"

type Course struct {
	name string
	price int
	url string
}

type Printer interface {
	printInfo() string
}

func (c * Course) printInfo() string {
	print("ssss")
	return "课程信息"
}

//func print(i interface{}) {
//	fmt.Printf("%v \n", i)
//}
func print(i interface{})  {
	//类型断言
	switch v:= i.(type) {
	case string:
		fmt.Printf("%s(字符串)\n", v)
	case int:
		fmt.Printf("%d(整数)\n", v)

	}
}

func main() {
	//空接口

	var i interface{}
	// 空接口类似python中的object类
	i = 20
	print(i)
	//i = Course{}

	// 可以把任何类型赋值给空接口

	i = []string{"a", "b", "c"}
	//fmt.Println(i)
	print(i)
	//空接口可以作为map的值
	var teacherInfo = make(map[string]interface{})
	teacherInfo["name"] = "jack"
	teacherInfo["age"] = 18
	teacherInfo["weight"] = 76.3
	teacherInfo["course"] = []string{"django", "scrapy", "sanic"}
	print(teacherInfo)
	for k,v := range teacherInfo{
		fmt.Println(k,v)
	}
	//c := &Course{}
	//c.printInfo()
	var c Printer = &Course{}
	c.printInfo()
}
