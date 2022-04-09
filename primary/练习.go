package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	var redball [6] int
	rand.Seed(time.Now().UnixNano())

	for i:=0; i<len(redball); i++{
		temp := rand.Intn(6) + 1
		for j:=0; j<i; j++{
			if redball[j] == temp{
				temp = rand.Intn(6) + 1
				j = -1
				continue

			}

		}
		redball[i] = temp
	}
	for i:=0; i<len(redball); i++{
		fmt.Print(redball[i])
	}


}
