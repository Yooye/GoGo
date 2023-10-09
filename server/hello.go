package main

import (
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"
	house_pb "server/proto/gen/go"
)

func main() {
	//1. 使用protobuf生成的go结构体，按需创建结构体数据实例
	house := house_pb.House{
		OwnerName: "三丰",
		Area:      99.8,
		Type:      1,
		MasterRoom: &house_pb.HouseRoom{
			RoomName: "主卧",
			RoomArea: 20.5,
		},
		OtherRooms: []*house_pb.HouseRoom{
			{
				RoomName: "次卧",
				RoomArea: 10,
			},
			{
				RoomName: "书房",
				RoomArea: 10,
			},
		},
		State: 1,
	}
	fmt.Println(&house)             //2. 通过地址，打印观察数据实例
	b, err := proto.Marshal(&house) //3.将数据包转为二进制数据流，可以让微服务之间高效传输数据
	if err != nil {
		panic(err)
	}
	fmt.Printf("%X\n", b) // 0A06E4B889E4B8B010011D9A99C742

	var house1 house_pb.House
	err = proto.Unmarshal(b, &house1) //4. 将二进制数据流，重新转回原始结构体格式
	if err != nil {
		panic(err)
	}
	fmt.Println(&house1)

	s, err := json.Marshal(&house) //5. 将结构体数据，转为json数据，可以提供给前端使用
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s) //{"owner_name":"三丰","type":1,"area":99.8}

}
