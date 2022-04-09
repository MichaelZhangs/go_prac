package main

import "fmt"

func main()  {
	var arr [5]int = [5]int{1,2,3,4,5}
	// 定义指针指向数组
	//fmt.Println(arr)
	fmt.Printf("%p\n", &arr)
	fmt.Printf("%p\n", &arr[0])
	//p := &arr
	var p *[5] int
	p = &arr
	fmt.Printf("%p\n", p)
	p[4] = 123
	fmt.Println(p[4])
	fmt.Println((*p)[4])
	fmt.Println(arr)
	fmt.Println(*p)
}
