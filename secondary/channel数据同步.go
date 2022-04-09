package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		for i:=0;i<2; i++{
			fmt.Println("i = ", i)
		}
		ch <- "finished"
	}()
	str := <- ch
	fmt.Printf(str)
}
