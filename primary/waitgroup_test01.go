package main

import (
	"fmt"
	"sync"
)
// 如何解决主线程goroutine 在子协程结束后自动结束
var wg sync.WaitGroup
// WaitGroup 提供三个很有用的函数
/*
	Add
	Done
	Wait
 */

func f(n int )  {
	defer wg.Done()
	fmt.Println(n)
}

func main() {

	for i:=0; i<5; i++{
		wg.Add(1)
		go f(i)
	}
	wg.Wait()
}
