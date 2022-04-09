package main

import (
	"fmt"
	"gopra/rpcPra/client_proxy"
)

func main() {
	//// 建立连接
	//client, _ := rpc.Dial("tcp", ":5555")
	////var reply *string = new(string)
	//var reply string
	//_ = client.Call("HelloService.Hello", "3333", &reply)
	////fmt.Println(*reply)
	//fmt.Println(reply)
//	基于json格式的rpc
//	conn, _ := net.Dial("tcp", "127.0.0.1:5555")
//	var reply string
//	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
//	_ = client.Call(handler.HelloServiceName+".Hello", "你是谁sdfg", &reply)
//
//	fmt.Println(reply)
	client := client_proxy.NewHelloServiceClient("127.0.0.1:5555", "tcp")
	var reply string
	err := client.Hello("哈ff哈哈", &reply)
	if err != nil{
		panic("调用错误...")
	}
	fmt.Println(reply)
}
