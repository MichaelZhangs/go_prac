package main

import (
	"fmt"
	"math/rand"
	"time"
)



func main() {
	a := rand.Intn(123)
	fmt.Println(a)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<10; i++{
		fmt.Println(rand.Intn(100))
	}
	fmt.Println(time.Now())
	fmt.Println(time.Now().Year())
}
