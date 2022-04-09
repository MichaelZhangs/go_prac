package main

import (
	"fmt"
	"net"
)

func main()  {
	conn, _ := net.Dial("tcp","127.0.0.1:8000")

	httpRequest := "GET /itcast HTTP/1.1\r\nHost:127.0.0.1:8000\r\n\r\n"

	conn.Write([]byte(httpRequest))
	defer conn.Close()
	buf := make([]byte, 4096)
	n, _ := conn.Read(buf)
	if n == 0{
		return
	}
	fmt.Println(string(buf[:n]))
}
