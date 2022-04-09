package main

import (
	"fmt"
	"net"
)
// 创建用户结构体
type Client struct {
	C chan string
	Name string
	Addr string
}

// 创建全局map, 存储在线用户
var onlineMap map[string]Client

// 创建全局channel, 传递用户消息
var message = make(chan string)

func Manager()  {
	// 初始化 onlineMap
	onlineMap = make(map[string]Client)
	for { // 循环从message 中读消息

		// 监听message
		msg := <- message
		// 循环发送消息给在线用户
		for _, clt := range onlineMap {
			clt.C <- msg
		}
	}
}

func writeMessage(client Client, conn net.Conn)  {
	// 监听用户 自带channel 上有消息
	for msg := range client.C{
		conn.Write([]byte(msg + "\n"))
	}
}

func makeMsg(client Client, msg string) (buf string) {
	buf = "[" + client.Addr +"] " +client.Name + " "+ msg
	return
}

func HandlerConnectChat(conn net.Conn){
	defer conn.Close()
	//
	// 获取用户的 网络地址
	cltAddr := conn.RemoteAddr().String()
	fmt.Println("cltAddr : ", cltAddr)
	// 创建新连接用户的 结构体
	clt := Client{make(chan string),cltAddr,cltAddr}
	//  将新连接用户添加到在线用户map中
	onlineMap[cltAddr] = clt
	// 创建专门用于发送消息的go 程
	go writeMessage(clt,conn)

	// 发送用户上线消息到 全局 channel 中
	message <- makeMsg(clt, " 上线了: ")

	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0{
				fmt.Println("检测到客户端 %s 退出 \n", clt.Name)
				return
			}
			if err != nil{
				fmt.Println("conn.ReadError")
				return
			}
			// 将读取的消息广播给在线用户
			msg := string(buf[:n])
			message <- makeMsg(clt, msg)
		}

	}()

	for  {
		;
	}
}
func main() {
	// 创建监听套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err!= nil{
		fmt.Println("net.listen err: ", err)
		return
	}
	defer listener.Close()

	// 创建管理者go 程
	go Manager()
	// 循环监听客户端请求
	for {
		conn, err := listener.Accept()
		if err != nil{
			fmt.Println("listener.accept err: ", err)
			return
		}
		// 启动go程处理客户端请求
		go HandlerConnectChat(conn)

	}
}
