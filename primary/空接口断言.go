package main

import "fmt"

func main(){

	arr := make([]interface{}, 3)
	arr = append(arr, 1,1.23,"Hello aaa")
	for _, v := range arr{
		if data, ok := v.(int); ok{
			fmt.Println("整型 ", data)
		}else if data, ok := v.(float64); ok{
			fmt.Println("浮点型",data)
		}else if data, ok := v.(string); ok{
			fmt.Println("字符串", data)
		}

	}

}
