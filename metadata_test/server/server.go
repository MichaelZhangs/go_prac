package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"gopra/grpcPra/proto"
	"net"
)

type Server struct {

}

func (s * Server )SayHello(ctx context.Context, request *proto.HelloRequest) (* proto.HelloReply, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	//fmt.Println("ok = ", ok, "md = ", md)
	if !ok {
		fmt.Println("get metadata error")
	}
	fmt.Println(" name = ", md.Get("name")[0])
	fmt.Println("pwd = ", md["pwd"][0])
	if nameSlice, ok := md["name"]; ok {
		fmt.Println(nameSlice)
		for i, e := range nameSlice{
			fmt.Println(i, e)
		}
	}

	return &proto.HelloReply{
		Message: "很爱很爱你 " + request.Name ,

	}, nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	lis , err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		panic("failed to listen")
	}
	g.Serve(lis)

}
