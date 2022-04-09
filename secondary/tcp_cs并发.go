package main

import (
	"fmt"
	"net"
	"strings"
)

func HandlerConnect(conn net.Conn){
	defer conn.Close()
	addr := conn.RemoteAddr()
	fmt.Println(addr, "客户端成功连接")

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.read err :", err)
			return
		}
		fmt.Println("收到的数据: ", string(buf[:n]))
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}

}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:5000")
	if err != nil{
		fmt.Println("net.listener err: ", err)
		return
	}
	defer listener.Close()

	for {
		fmt.Println("服务器等待连接。。。。。")
		conn, err:= listener.Accept()
		if err != nil {
			fmt.Println("net.listener err: ", err)
			return
		}
		// 具体完成服务器和客户端的通信
		go HandlerConnect(conn)
	}
}
