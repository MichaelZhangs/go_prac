package main

import (
	"fmt"
	"io"
	"net/http"
)

func main()  {
	resp, err := http.Get("http://www.baidu.com")

	if err != nil{
		fmt.Println("Get err: ", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Header: ", resp.Header)
	fmt.Println("status: ", resp.Status)
	fmt.Println("StatusCode: ", resp.StatusCode)
	fmt.Println("proto: ", resp.Proto)
	fmt.Printf("")

	buf := make([]byte, 4096)
	var result string

	for {
		n, err := resp.Body.Read(buf)
		if n == 0{
			fmt.Println("Read finished....")
			return
		}
		if err != nil && err != io.EOF{
			fmt.Println("err : ", err)
			return
		}
		result += string(buf[:n])
	}
	fmt.Printf("|%v|\n", result)
}
