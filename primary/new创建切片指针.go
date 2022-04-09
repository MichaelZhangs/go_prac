package main

import "fmt"

func main(){
	var p *[]int
	fmt.Printf("%p\n", p)
	p = new([]int)
	fmt.Printf("%p\n", p)
	fmt.Println(p)
	*p = append(*p, 1,2,3,4,5)
	fmt.Printf("%p\n", p)
	fmt.Println(p)
	for i:=0; i<len(*p); i++ {
		fmt.Println((*p)[i])
	}
	for k,v := range *p{
		fmt.Println(k, v)
	}
}
