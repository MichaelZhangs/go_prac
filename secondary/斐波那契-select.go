package main

import "fmt"

func fib(ch <- chan int, quit <- chan bool){
	for {
		select {
			case num := <- ch:
				fmt.Print(num, " ")
			case <- quit:
				return
		}
	}
}

func main() {
	ch := make(chan int )
	quit := make(chan bool)
	go fib(ch, quit) // 打印数列
	x, y := 1, 1
	for i:=0; i<20; i++{
		ch <- x
		x, y = y, x+y
	}
	quit <- true
}
