package main

import (
	"context"
	"fmt"
	house_pb "server/proto/gen/go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	connect, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer connect.Close()
	c := house_pb.NewHouseServiceClient(connect)
	//1. create metadata
	md := metadata.New(map[string]string{
		"auth": "123",
	})
	//2. create a new context with some metadata
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	//3. send request within metadata
	r, err := c.GetHouse(ctx, &house_pb.GetHouseRequest{Id: "111"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}
