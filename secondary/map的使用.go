package main

import "fmt"

func main(){
	// 方法一 创建
	m := make(map[int]string)
	m[1] = "222"
	m[2] = "3"
	fmt.Println(m)
	fmt.Println(len(m))
//	 方法二 创建
	m2 := map[int]string{}
	m2[1] = "a"
	fmt.Println(m2)
	delete(m,1)
	var a string
	fmt.Scan(&a)
	fmt.Println("a = ", a)
}
