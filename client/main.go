package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "memo-RPC/client/ecommerce"
)

const PORT = "9001"

func main() {
	// 拨号连接
	conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()
	// 创建新客户端，使用拨号创建的连接
	client := pb.NewEventServiceClient(conn)

	// 设置过期时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 访问该函数，获取 响应
	resp, err := client.ShowEvent(ctx, &pb.ShowEventRequest{
		Id: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("Greeting: %s", resp.String())
}
