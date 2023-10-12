package main

import (
	"context"
	"fmt"
	house_pb "server/proto/gen/go"

	"google.golang.org/grpc"
)

type customCredential struct {
}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"auth": "1234",
	}, nil
}
func (c customCredential) RequireTransportSecurity() bool {
	return false
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithPerRPCCredentials(customCredential{}))
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
