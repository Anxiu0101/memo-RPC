package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "memo-RPC/client/ecommerce"
)

const UserPORT = "9002"
const EventPORT = "9001"

func main() {
	go testUserService()
	go testEventService()
}

func testUserService() {
	// 拨号连接
	conn, err := grpc.Dial(":"+UserPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()
	// 创建新客户端，使用拨号创建的连接
	client := pb.NewUserServiceClient(conn)

	// 设置过期时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 访问该函数，获取 响应
	resp1, err := client.Register(ctx, &pb.UserRegisterRequest{
		Username: "Anxiu",
		Password: "123456",
	})
	if err != nil {
		log.Fatalf("client.Register err: %v", err)
	}
	log.Printf("Greeting: %s", resp1.String())

	resp2, err := client.Login(ctx, &pb.UserLoginRequest{
		Username: "Anxiu",
		Password: "123456",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}
	log.Printf("Token: %s", resp2.Token)
}

func testEventService() {
	// 拨号连接
	conn, err := grpc.Dial(":"+EventPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
