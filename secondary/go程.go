package main

import (
	"fmt"
	"runtime"
)

func main()  {
	go func() {
		for i:=0; i< 5; i++{
			fmt.Println("hello")
			//time.Sleep(time.Millisecond)
		}
	}()
	go func() {
		for i:=0; i< 5; i++{
			fmt.Println("world")
			runtime.Goexit()
			//time.Sleep(time.Second)
		}
	}()
	for i:=0; i< 3; i++{
		runtime.Gosched()
		fmt.Println("lkkk")

	}

}
