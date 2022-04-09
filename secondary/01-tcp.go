package main

import (
	"fmt"
	"net"
)

func main()  {
	// 指定服务器通信协议，ip地址和端口号
	listener, err := net.Listen("tcp", "127.0.0.1:5000")
	if err != nil {
		fmt.Println("net.listener err :", err)
		return
	}
	defer listener.Close()
	fmt.Println("等待建立连接.....")
	// 阻塞监听客户端连接请求
	conn , err := listener.Accept()
	defer conn.Close()
	if err != nil{
		fmt.Println("listener.accept err :", err)
		return
	}
	fmt.Println("服务器与客户端建立连接")
	// 读取客户端发送的数据
	buff := make([]byte, 1024)
	n, err := conn.Read(buff)
	if err != nil{
		fmt.Println("conn.Read err :", err)
		return
	}
	// 处理数据
	fmt.Println(" 服务器读取的数据: ", string(buff[:n]))
	conn.Write([]byte("hehehehehe"))


}
