package main

import (
	"fmt"
	"time"
)

func main() {
	myTicker := time.NewTicker(time.Second)
	go func() {
		nowTime := <- myTicker.C
		fmt.Println(nowTime)
	}()
	for {;}
}
