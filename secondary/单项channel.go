package main

import "fmt"

func send(sendCh chan <- int){
		sendCh <- 989
		close(sendCh)
}
func recv(in  <- chan int ){
	num := <- in
	fmt.Println(num)
}

func main() {
	ch := make(chan  int ) // 双向
	//go func() {
	//	send(ch)
	//}()
	go send(ch)
	recv(ch)
}
