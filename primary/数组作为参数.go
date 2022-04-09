package main

import "fmt"

func test1(arr [10] int) {
	for i:=0; i<len(arr); i++{
		arr[i] = arr[i] + 2
		fmt.Println(arr[i])
	}
	arr[9] = 100
	fmt.Printf("%p\n", &arr)
}

func main() {
	// 指定数组作为初始化参数 6:10 表示 默认为0
	var arr [10] int = [10] int{1,2,3,4,5,6:10}
	fmt.Println(arr)
	test1(arr)
	fmt.Println(arr)
	fmt.Printf("%p\n", &arr)

}