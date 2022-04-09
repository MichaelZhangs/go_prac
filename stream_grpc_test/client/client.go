package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopra/stream_grpc_test/proto"
	"sync"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8866", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	if err != nil{
		panic(err)
	}
	// 服务端流模式
	c := proto.NewGreeterClient(conn)
	res, _ :=c.GetStream(context.Background(), &proto.StreamReqData{Data: "慕课网"})
	for {
		a, err := res.Recv()
		if err != nil{
			fmt.Println(err)
			break
		}
		fmt.Println(a.String())
	}
	// 客户端流模式
	putS, _ := c.PutStream(context.Background())
	i := 0
	for {
		i ++
		_ = putS.Send(&proto.StreamReqData{
			Data: fmt.Sprintf("慕课网 %d", i),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	// 双向流模式
	allStr, _ := c.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			data , _ := allStr.Recv()
			fmt.Println("收到服务端消息: "+ data.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto.StreamReqData{Data: "我是慕课网"})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()


}


