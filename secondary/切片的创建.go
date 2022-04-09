package main

import (
	"fmt"
)


func noEmpty(data []string) []string{
	out := data[:0]
	for _, v := range(data){
		if v != ""{
			out = append(out, v)
		}
	}
	return out
}


func main(){

	arr := []string{"red", "","black", "","", "yellow","red", "","black", "","", "yellow"}
	data := noEmpty(arr)
	a := arr[:2]
	fmt.Println(arr)
	fmt.Println(data)
	copy(a, arr[:4])
	fmt.Println(a)
}
