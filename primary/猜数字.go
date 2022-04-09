package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	rand.Seed(time.Now().UnixNano())

	temp := rand.Intn(900) + 100
	var num int
	for{
		fmt.Printf("请输入一个三位数 ")
		fmt.Scan(&num)
		if num > 99 && num < 1000{
			if num > temp{
				fmt.Println("输入数字过大")
			}else if num < temp{
				fmt.Println("输入数字过小")
			}else {
				fmt.Println("输入数字相等")
				break
			}
		}else{
			fmt.Println("输入错误， 重新输入")
		}
	}
	fmt.Println("随机数: ", temp)
}
