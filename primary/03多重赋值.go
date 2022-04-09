package main

import "fmt"

func main(){

	a, b, c, d := 10, 20, 30.5, 40
	a, b = b, a
	fmt.Println(a, b, c, d)
}
