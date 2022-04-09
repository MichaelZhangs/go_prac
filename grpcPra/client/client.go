package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	"gopra/grpcPra/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		panic(err)
	}
	defer  conn.Close()
	c := proto.NewGreeterClient(conn)
	r , err := c.SayHello(context.Background(), &proto.HelloRequest{
		Name: "jack",

		AddTime:timestamppb.Now(),

	})
	if err != nil{
		panic(err)
	}
	fmt.Println(r.Message)

}
