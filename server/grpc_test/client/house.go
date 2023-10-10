package main

import (
	"context"
	"fmt"
	house_pb "server/proto/gen/go"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

func main() {
	// 1. create interceptor function
	interceptor := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		fmt.Printf("客户端请求拦截器\n")
		md := metadata.New(map[string]string{
			"auth": "1234",
		})
		ctx = metadata.NewOutgoingContext(context.Background(), md)
		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	}
	// 2. create a client Interceptor
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	// 3. set the Interceptor to grpc
	connect, err := grpc.Dial("127.0.0.1:8088", opts...)
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
