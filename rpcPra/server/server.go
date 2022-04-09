package main

import (
	"gopra/rpcPra/handler"
	"gopra/rpcPra/server_proxy"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)


func main() {
	// 实例化 server
	listener, _ := net.Listen("tcp", "127.0.0.1:5555")
	// 注册处理逻辑
	//_ = rpc.RegisterName(handler.HelloServiceName, &handler.HelloService{})
	_ = server_proxy.RegisterHelloService(&handler.HelloService{})
	// 启动服务
	for  {
		conn, _ := listener.Accept()
		//rpc.ServeConn(conn)
	//	基于json
	 	go  rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}


}
