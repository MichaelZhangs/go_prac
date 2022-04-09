package client_proxy

import (
	"gopra/rpcPra/handler"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloServiceStub struct {
	* rpc.Client
}
// 初始化
func NewHelloServiceClient(address, proto string) HelloServiceStub {
	conn, err := net.Dial(proto, address)

	if err != nil{
		panic("connect error")
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return HelloServiceStub{client}
}

func (c * HelloServiceStub)Hello(request string, reply * string)  error {
	err := c.Call(handler.HelloServiceName+".Hello", request, &reply)
	if err != nil{
		return err
	}
	return nil
}

