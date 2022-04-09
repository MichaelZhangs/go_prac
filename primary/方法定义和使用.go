package main

import "fmt"

//func add(a, b int ) int {
//	return a+b
//}
//方法的定义和使用

type Int int

// func (方法接收者)方法名(参数列表) 返回值类型
func (a Int) add(b Int) Int{
	return a+b
}

func main(){
	//s := add(10,20)
	//fmt.Println(s)
	var a Int = 30
	//var b Int
	s := a.add(20)
	fmt.Println(s)
}
