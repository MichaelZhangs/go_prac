package main

import "fmt"

func main()  {
	fmt.Println("111111")
	// defer 延迟调用
	defer fmt.Println("222222")
	defer fmt.Println("333333")
	fmt.Println("444444")

	x := 10
	defer func(a int ) {
		x ++
		fmt.Println("defer ", x)
	}(x)
	fmt.Println(x)

}
