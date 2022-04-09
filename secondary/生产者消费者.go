package main

import "fmt"

type OrderInfo struct {
	id int
}

func producer(out chan <- OrderInfo){
	for i:=0; i < 10;i++{
		order := OrderInfo{id: i*i}
		out <- order
	}
	close(out)
}
func consumer(in <- chan OrderInfo){

	for order := range in{
		fmt.Println("消费的id : ", order.id)
	}
}

func main()  {

	ch := make(chan OrderInfo)
	go producer(ch)
	consumer(ch)
}
