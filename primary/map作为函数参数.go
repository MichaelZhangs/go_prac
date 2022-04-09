package main

import "fmt"

func test(m map[int]string) map[int]string{
	m[3] = "c"
	return m
}

func main(){
	m := map[int]string{1:"a",2: "b"}
	// map 作为参数传递是引用传递
	s := test(m)
	fmt.Println(m)
	fmt.Println(s)
	fmt.Printf("%p %p", m, s)
}
