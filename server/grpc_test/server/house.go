package main

import (
	"context"
	"fmt"
	"net"
	house_pb "server/proto/gen/go"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc/metadata"

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
		Msg: "auth验证成功,并下发了数据",
	}, nil
}

func main() {
	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return resp, status.Error(codes.Unauthenticated, "未携带auth信息")
		}
		fmt.Printf("服务端拦截器", md)
		if md["auth"][0] != "123" {
			return resp, status.Error(codes.Unauthenticated, "auth错误")
		}
		return handler(ctx, req)
	}
	opt := grpc.UnaryInterceptor(interceptor)
	g := grpc.NewServer(opt)
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
