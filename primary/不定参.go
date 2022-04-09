package main

import (
	"fmt"
)

func test(args ...int) []int {
	fmt.Println(args)
	return args
}

func main(){
	a := test(1,2,3)
	fmt.Println(len(a))
	fmt.Println(a[0])
	for i, data := range a{
		fmt.Println(i, data)

	}
}
