package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	house_pb "server/proto/gen/go"
)

func main() {
	connect, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer connect.Close()
	c := house_pb.NewHouseServiceClient(connect)
	r, err := c.GetHouse(context.Background(), &house_pb.GetHouseRequest{Id: "111"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}
