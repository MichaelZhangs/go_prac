package main

import "fmt"

func main01(){
	var a int = 10

	var p * int = &a
	fmt.Println(p)
	fmt.Println(*p)
	*p = 120
	fmt.Println(a)

}
func swap1(a, b int) ( int , int ){
	a, b = b, a
	fmt.Println("a = ", a, "b = ", b)
	return a, b
}
func swap2(x, y *int){
	*x, *y = *y, *x
	fmt.Println("x = ", &x, "y = ", &y)
}

func main(){

	a, b := 10, 20
	//a, b = swap1(a, b)
	fmt.Println("a = ", a, "b = ", b)
	swap2(&a, &b)
	fmt.Println("a = ", &a, "b = ", &b)
	fmt.Println("a = ", a, "b = ", b)
}