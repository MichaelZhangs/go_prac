package main

import (
	"errors"
	"fmt"
)

func div(a, b int ) ( res int ,err error) {

	if b == 0{
		err = errors.New("被除数不为0")
	}else {
		res = a / b
	}
	return
}
func main() {
	res, err := div(12, 4)
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println(res)
	}


}
