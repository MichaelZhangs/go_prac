package main

import (
	"fmt"
	"net"
)


func main()  {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	defer listener.Close()
	if err != nil{
		fmt.Println("net.listener err: ", err)
		return
	}
	conn, err := listener.Accept()
	defer conn.Close()
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if n == 0 {
		return
	}
	fmt.Println(string(buf[:n]))
}

