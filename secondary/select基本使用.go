package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan  int )
	quit := make(chan bool )

	go func() {
		for i:=0; i<5 ; i++{
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
		quit <- true  //通知主go 程推出
		runtime.Goexit()

	}()
	for  {
		select {
		case num := <- ch:
			fmt.Println("num = ", num)
		case <- quit:
			//break // 跳出select
			return  // 终止进程
		}
		fmt.Println("======")
	}
}
