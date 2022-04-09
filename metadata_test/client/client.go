package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"gopra/grpcPra/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		panic(err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)

	//md := metadata.Pairs("timestamp", time.Now().Format("2006-10-24 03:04:12"))
	md := metadata.New(map[string]string{
		"name": "jjjj",
		"pwd": "j123456",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r , err := c.SayHello(ctx, &proto.HelloRequest{
		Name: "jack",

	})
	if err != nil{
		panic(err)
	}
	fmt.Println(r.Message)

}


