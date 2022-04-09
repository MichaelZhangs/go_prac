package main

import "fmt"

func test(n []int){
	n[3] = 111
	fmt.Println(n)
	fmt.Printf("%p\n", n)
}

func main(){
	s := []int{1,2,3,4,5}
	test(s)
	fmt.Println(s)
	fmt.Printf("%p\n", s)
}
