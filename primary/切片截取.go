package main

import "fmt"

func main(){
	s := [] int{1,2,3,4,5}
	slice:= s[ 2:4:4]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Printf("%p %p\n", s, slice)
	slice[1] = 111
	fmt.Println(slice)
	fmt.Println(s)
	fmt.Println(cap(slice))
	s = append(s, 6,7,8)
	fmt.Println(slice)
	fmt.Println(s)
	fmt.Printf("%p %p\n", s, slice)
}
