package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cpuInfo(ctx context.Context){
		defer wg.Done()
		//ctx2, _ := context.WithCancel(ctx)
		//go memInfo(ctx2)
		for {
			select {
			case <- ctx.Done():
				fmt.Println("监控退出")
				return
			default:
				time.Sleep(time.Second)
				 fmt.Println("获取cpu信息成功")
			}
		}
}

func memInfo(ctx context.Context)  {
	defer wg.Done()

	for {
		select {
		case <- ctx.Done():
			fmt.Println("监控内存退出")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("获取内存信息成功")
		}
	}
}

func main() {
	/*
	 go 语言提供了一个select 的功能，作用于 channel之上的
	 io 多路复用
	select 随机公平的取数据
	 */
	//ch := make(chan int , 1)
	//ch2 := make(chan int , 1)
	//
	//ch <- 1
	//ch2 <- 2
	//
	//select {
	//case data := <- ch:
	//	fmt.Println(data)
	//case data := <- ch2:
	//	fmt.Println(data)
	//}

	wg.Add(1)
	//ctx , cancel := context.WithCancel(context.Background())
	ctx , _ := context.WithTimeout(context.Background(), time.Second * 3)
	go cpuInfo(ctx)
	//time.Sleep(time.Second * 6)
	////fmt.Println("cancel = ", cancel)
	//cancel()
	wg.Wait()
}
