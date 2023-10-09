package main

import (
	"context"
	"fmt"
	"net"
	house_pb "server/proto/gen/go"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

type Service struct {
	house_pb.HouseServiceServer
}

func (*Service) GetHouse(ctx context.Context, req *house_pb.GetHouseRequest) (*house_pb.GetHouseResponse, error) {
	// 1. get the metadata from context
	md, ok := metadata.FromIncomingContext(ctx)
	data := &house_pb.House{}
	msg := "auth验证失败"
	if ok {
		fmt.Println("收到了metadata", md["auth"])
		//2. use metadata to auth something
		if md["auth"][0] == "123" {
			data = &house_pb.House{
				OwnerName: "三丰",
				Area:      99.8,
				Type:      1,
				MasterRoom: &house_pb.HouseRoom{
					RoomName: "主卧",
					RoomArea: 20.5,
				},
				State: 1,
			}
			msg = "auth验证成功"
		}
	}
	//3. send authed data to client
	return &house_pb.GetHouseResponse{
		Id:   req.Id,
		Data: data,
		Msg:  msg,
	}, nil
}

func main() {
	g := grpc.NewServer()
	house_pb.RegisterHouseServiceServer(g, &Service{})
	lis, err := net.Listen("tcp", "0.0.0.0:8088")
	if err != nil {
		panic("端口监听失败")
	}
	err = g.Serve(lis)
	if err != nil {
		panic("rgpc服务启动失败")
	}
}
