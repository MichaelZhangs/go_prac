package main

import "fmt"

func varname(){
	PI := 3.14159  // 定义变量 自动推到类型
	r := 2.5
	// 在go语言中，不同类型不能计算, 可以通过类型转换
	s := PI * r * r   // 定义推倒类型
	l := PI * r * 2
	fmt.Println("面积为: ", s)
	fmt.Println("周长为: ", l)
}

func main()  {
	var a int
	a = 10
	fmt.Println("a = ", a)
	varname()
	s_even := 0
	s_odd := 0
	for i:= 1; i <= 100; i++ {
		if i%2 == 0 {
			s_even += i
		}else {
			s_odd += i
		}
	}
	fmt.Println("偶数和是: ", s_even)
	fmt.Println("奇数和是: ", s_odd)
}
