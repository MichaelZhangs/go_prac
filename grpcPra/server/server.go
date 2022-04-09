package main

import (
	"context"
	"google.golang.org/grpc"
	"gopra/grpcPra/proto"
	"net"
)

type Server struct {

}

func (s * Server )SayHello(ctx context.Context, request *proto.HelloRequest) (* proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "很爱很爱你 "+request.Name + " \n " + request.AddTime.String(),

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