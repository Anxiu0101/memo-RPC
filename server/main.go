package main

import (
	"log"
	"memo-RPC/server/conf"
	"memo-RPC/server/ecommerce"
	"memo-RPC/server/model"
	"memo-RPC/server/service"
	"net"

	"google.golang.org/grpc"
)

func init() {
	conf.Setup()
	model.Setup()
}

const PORT = ":9001"

func main() {

	// 监听端口
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	// 创建新 grpc 服务
	server := grpc.NewServer()

	// 将服务注册到服务端
	ecommerce.RegisterEventServiceServer(server, &service.EventService{})
	ecommerce.RegisterUserServiceServer(server, &service.UserService{})
	log.Printf("server listening at %v", lis.Addr())

	// 调用服务
	if err := server.Serve(lis); err != nil {
		log.Fatalf("fail to server: %v", err)
	}
}
