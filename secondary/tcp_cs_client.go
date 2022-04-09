package main

import (
	"fmt"
	"net"
)

func main() {
	// 指定服务器的ip
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil{
		fmt.Println("net.Dial err: ", err)
		return
	}
	defer conn.Close()
	// 主动写
	conn.Write([]byte("are u ok!!!"))
	// 接受服务器回来的数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	defer conn.Close()
	if err != nil{
		fmt.Println("conn.Read err ", err)
		return
	}
	fmt.Println("服务器的数据: ", string(buf[:n]))
	for  {
		;
	}
}
