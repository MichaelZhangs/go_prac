package main

import (
	"fmt"
	"time"
)

func pointer(s string)  {
	for _, ch := range s{
		fmt.Printf("%c",ch)
		time.Sleep(300 * time.Millisecond)
	}
}


func main(){

	ch := make(chan []string)
	fmt.Println(ch)



}
