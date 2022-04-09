package main

import "fmt"

func main(){
	arr := []int{1,2,3,4,5}
	fmt.Println(arr)
	s := arr[2:4]
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println(cap(arr))
	s2 := s[1:3]
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

}
