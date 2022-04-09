package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	conn, _ := net.Dial("udp", "127.0.0.1:6000")
	defer conn.Close()
	fmt.Printf("输入: ")
	buf := make([]byte, 1024)
	n ,_ := os.Stdin.Read(buf)
	conn.Write([]byte(buf[:n]))
	defer conn.Close()
	n, _ = conn.Read(buf)
	fmt.Println("接受: ",string(buf[:n]))
}
