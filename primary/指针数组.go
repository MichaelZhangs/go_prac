package main

import "fmt"

func main(){
	// 指针数组 --> 数组存储的是指针
	var arr [3]*int
	a := 10
	b := 20
	c := 30
	//arr[0] = &a
	//arr[1] = &b
	//arr[2] = &c
	//arr = append(arr , &a, &b, &c)   arr 是切片
	fmt.Println(arr)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", &c)

}
