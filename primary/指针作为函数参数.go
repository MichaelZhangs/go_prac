package main

import "fmt"

func swap(a int, b int){

	a , b = b, a
}
// ζιδΌ ι
func swap2(a * int, b *int){
	*a, *b = *b, *a

}
func main(){

	a := 10
	b := 20
	swap(a, b)
	fmt.Println(a, b)
	swap2(&a, &b)
	fmt.Println(a, b)
}
