package main

import (
	"context"
	"net"
	house_pb "server/proto/gen/go"

	"google.golang.org/grpc"
)

type Service struct {
	house_pb.HouseServiceServer
}

func (*Service) GetHouse(ctx context.Context, req *house_pb.GetHouseRequest) (*house_pb.GetHouseResponse, error) {
	return &house_pb.GetHouseResponse{
		Id: req.Id,
		Data: &house_pb.House{
			OwnerName: "三丰",
			Area:      99.8,
			Type:      1,
			MasterRoom: &house_pb.HouseRoom{
				RoomName: "主卧",
				RoomArea: 20.5,
			},
			State: 1,
		},
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
