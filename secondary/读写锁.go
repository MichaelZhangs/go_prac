package main

import (
	"fmt"
	"math/rand"
	"time"
)

func read(in <- chan int , a int )  {
	for {
		num := <- in
		fmt.Printf("%d 读go , 读入 %d\n", a , num)
	}

}
func write(out  chan <- int, b int  ) {
	for {
		// 生成随机数
		num := rand.Intn(1000)
		out <- num
		fmt.Printf("%d 写go , 写入 %d\n", b, num)
		time.Sleep(time.Millisecond * 200)
	}

}

func main() {
	//
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int)
	//quit := make(chan bool)

	for i:=0; i<5; i++{
		go read(ch, i)


	}
	for i:=0; i<5; i++{
		go write(ch, i)

	}
	for {
		;
	}
	//<- quit
}
