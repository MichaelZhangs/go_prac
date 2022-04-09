package main

import "fmt"


func test(a []string) []string{
	b := append(a, "4,5,6")
	return b
}

func main() {
	name := "jack舒服撒"

	for _, v := range name{
		fmt.Printf("%c", v)
	}
	fmt.Println(name[6])
	name_arr := []rune(name)

	// 如下打印中文不会乱码
	fmt.Println(name_arr)
	for i:=0; i<len(name_arr); i++{
		fmt.Printf("%c", name_arr[i])
	}

	var a = []string{}
	a = append(a, "a", "b", "c")
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	a = append(a, "1", "2")
	fmt.Println(a)
	a = test(a)
	fmt.Println(a)

}
