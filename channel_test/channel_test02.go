package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func consumer(q chan int ){
	defer wg.Done()
	fmt.Println(<-q)
}

func main() {
	var msg chan int
	msg = make(chan int )
	wg.Add(1)
	go consumer(msg)
	msg <- 1

	close(msg)
	wg.Wait()
	//data := <- msg
	//fmt.Println(data)
}
