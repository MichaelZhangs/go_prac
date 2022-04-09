package main

import (
	"fmt"
	"strings"
)

func test(s []string) [] string  {
	s = append(s, "a", "b")
	return s
}

func main() {
	// 切片设置
	s := make([]string, 5)
	s = append(s,  "1", "2", "3", "4","5", "6", "7")
	fmt.Println(s[1])
	s = test(s)
	fmt.Println(s)
	a := []string{"c", "d"}
	b := append(s, a...)
	fmt.Println(b)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	dic := make(map[string]int, 10)
	fmt.Println(dic)
	dic["1"] = 2
	fmt.Println(dic)
	s1 := "a-b-c-d"
	s_arr := strings.Split(s1, "-")
	fmt.Println(s_arr)
	s2 := strings.Join(s, "-")
	fmt.Println(s2)
}
