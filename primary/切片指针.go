package main

import "fmt"

func main(){
	var arr []int = []int{1,2,3,4,5}

	p := &arr
	(*p)[1] = 222
	fmt.Printf("%p\n", &arr)
	fmt.Printf("%p\n", &arr[0])
	fmt.Printf("%p\n", p)
	fmt.Printf("%p\n", arr)
	//fmt.Printf("%T\n", &arr)
	//fmt.Printf("%T\n", p)
	var k *[]int
	k = &arr
	(*k)[2] = 333   // 切片使用方式 切片本身就是歌地址， 而数组不是

	*k = append(*k, 6,7,8,9)
	fmt.Printf("%p\n", arr)
	fmt.Printf("%p\n", *k)
	//k := arr
	//fmt.Printf("%p\n", k)
	fmt.Println(arr)
}
