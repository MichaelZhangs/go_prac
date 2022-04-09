package main

import "fmt"

// 匿名函数， 实现函数在栈区的本地化操作
func test2() func() int{
	var a int
	return func() int{
		a ++
		return a
	}
}


func main(){
	f:= test2()
	fmt.Printf("%T", f)
	for i:=0; i<10; i++{
		fmt.Println(f())
	}
}