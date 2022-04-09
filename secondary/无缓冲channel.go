package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)

	go func() {
		for i:=0; i< 5; i++{
			//fmt.Println("i = ", i)
			ch <- i*i
			fmt.Println("i = ", i)
		}
		close(ch)
	}()
	time.Sleep(time.Second*2)
	//for i:=0; i<5; i++{
	//	//num := <- ch
	//
	//	if num, ok := <- ch; ok == true{
	//		fmt.Println("num = ", num)
	//	}
	//}
	for num:= range ch{
		fmt.Println( "num = ", num)
	}
}
