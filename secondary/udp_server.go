package main

import (
	"fmt"
	"net"
	"strings"
)


func main() {
	strAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:6000")
	if err != nil{
		fmt.Println("net.resolve err: ", err)
		return
	}
	fmt.Println("udp 服务器地址结构: ", strAddr)
	udpConn, err := net.ListenUDP("udp",strAddr)
	if err != nil{
		fmt.Println("net.ListenUDP err: ", err)
		return
	}
	defer udpConn.Close()
	// 读取客户端发送的数据
	buf := make([]byte, 1024)
	n, clientAddr ,err := udpConn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("udpConn.read err: ", err)
		return
	}
	defer udpConn.Close()
	fmt.Println("from ", clientAddr)
	fmt.Println("recv : ", string(buf[:n]))
	_, err = udpConn.WriteToUDP([]byte(strings.ToUpper(string(buf[:n]))), clientAddr)
	if err != nil {
		fmt.Println("udpConn.read err: ", err)
		return
	}
}
