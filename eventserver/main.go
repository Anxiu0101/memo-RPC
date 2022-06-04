package main

import (
	"context"
	"log"
	"memo-RPC/eventserver/conf"
	"memo-RPC/eventserver/ecommerce"
	"memo-RPC/eventserver/model"
	"memo-RPC/eventserver/service"
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

	// 使用一元拦截器（grpc.UnaryInterceptor），验证请求
	// TODO 增加流式请求拦截器
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// 拦截普通方法请求，验证Token
		err = service.CheckAuthority(ctx)
		if err != nil {
			return
		}
		// 继续处理请求
		return handler(ctx, req)
	}

	// 创建新 grpc 服务
	server := grpc.NewServer(grpc.UnaryInterceptor(interceptor))

	// 将服务注册到服务端
	ecommerce.RegisterEventServiceServer(server, &service.EventService{})
	log.Printf("server listening at %v", lis.Addr())

	// 调用服务
	if err := server.Serve(lis); err != nil {
		log.Fatalf("fail to server: %v", err)
	}
}