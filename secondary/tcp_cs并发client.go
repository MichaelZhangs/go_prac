package main

import (
	"fmt"
	"net"
	"os"
)

func HandlerConnectClinet(conn net.Conn)  {

	defer conn.Close()
	//var a string
	buf := make([]byte, 1024)

	for {
		fmt.Printf("输入: ", )
		//fmt.Scan(&a)
		n, err := os.Stdin.Read(buf)
		if err != nil{
			fmt.Println("stdin.Read err ", err)
			continue
		}
		conn.Write([]byte(string(buf[:n])))

		n, err = conn.Read(buf)
		if err != nil{
			fmt.Println("conn.Read err ", err)
			return
		}
		fmt.Println("接受到的数据 ", string(buf[:n]))
	}


}

func main()  {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil{
		fmt.Println("net.Dial err: ", err)
		return
	}
	defer conn.Close()
	for {
		HandlerConnectClinet(conn)
	}

}
