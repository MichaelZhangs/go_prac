package main

import "fmt"

func main(){
	// map[keyType] valueType
	//var m map[int]string
	m := make(map[int]string, 0)
	m[1] = "a"
	m[5] = "b"
	m[7] = "abc"
	for k,v := range m{
		fmt.Println(k,v)
	}
	fmt.Println(m[1])

	y := map[string]int{"a": 1, "b": 2, "c":3}
	fmt.Println(y)
	v,b := m[1]
	fmt.Println(v,b)
	// 删除map中的元素
	delete(m, 1)
	fmt.Println(m)
}
