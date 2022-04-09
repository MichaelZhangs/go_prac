package main

import (
	"fmt"
)

//func main(){
//	a := 10
//	b := 20
//	// 匿名函数  返回值形式 2
//	f := func (a int, b int) int{
//		fmt.Println(a, b, a+b)
//		return a*b
//	}
//	fmt.Println(f(a, b))
//}
//func main(){
//	a := 10
//	b := 20
//	// 匿名函数 1
//	f := func (a int, b int){
//		fmt.Println(a, b, a+b)
//
//	}
//	f(a, b)
//}
func main(){
		a := 10
		b := 20
		// 匿名函数
		f := func (a int, b int) ( x int , y int, c int){
			fmt.Println(a, b, a+b)
			return a, b, a*b
		}
		fmt.Println(f(a, b))
	}