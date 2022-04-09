package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func consumer(q chan int ){
	defer wg.Done()
	//fmt.Println("data = ", <- q)
	//for data := range q {
	//	fmt.Println("data = ", data)
	//}
	for {
		data, ok := <- q
		if !ok{
			fmt.Println(" ok = ", ok)
			break
		}
		fmt.Println(" ok = ", ok)
		fmt.Println("data = ", data)
	}
}


func main()  {
	/**
	channel 提供了一种通信机制， 定向， 消息队列
	 */
	var msg chan int

	// 初始化 使用make 的三种类型 slice map channel
	// channel 的make 初始化方式:
	msg = make(chan int) // 无缓冲空间
	msg = make(chan int , 1) // 有缓冲空间

	if msg == nil{
		fmt.Println("msg 为空")
	}
	// 缓冲空间满了， 会出现deadlock
	msg <- 100
	wg.Add(1)

	go consumer(msg)
	msg <- 200
	msg <- 300
	close(msg)

	wg.Wait()
}