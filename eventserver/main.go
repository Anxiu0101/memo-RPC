package main

import (
	"context"
	"google.golang.org/grpc/credentials"
	"log"
	"memo-RPC/eventserver/conf"
	pb "memo-RPC/eventserver/ecommerce"
	"memo-RPC/eventserver/model"
	"memo-RPC/eventserver/service"
	"net"

	"google.golang.org/grpc"
)

func init() {
	conf.Setup()
	model.Setup()
}

const PORT = ":9002"

func main() {

	// 监听本地端口
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	var opts []grpc.ServerOption

	// 使用证书文件和密钥文件为服务端构造 TLS 凭证
	certs, err := credentials.NewServerTLSFromFile("../certs/memo.pem", "../certs/memo.key")
	if err != nil {
		log.Printf("Failed to generate credentials %v", err)
	} else {
		opts = append(opts, grpc.Creds(certs))
	}

	// 使用一元拦截器（grpc.UnaryInterceptor），验证请求
	// TODO 增加流式请求拦截器
	var interceptor grpc.UnaryServerInterceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// 拦截普通方法请求，验证 Token
		log.Println("In Authority")
		log.Println("filter:", info)
		username, err := service.CheckAuthority(ctx)
		if err != nil {
			log.Fatalf("interceptor err: %v", err)
			return nil, err
		} else {
			log.Printf("%v pass Authority", username)
		}
		// 继续处理请求
		return handler(ctx, req)
	}

	opts = append(opts, grpc.UnaryInterceptor(interceptor))

	// 创建新 grpc 服务
	server := grpc.NewServer(opts...)

	// 将服务注册到服务端
	pb.RegisterEventServiceServer(server, &service.EventService{})
	log.Printf("Event server listening at %v", lis.Addr())

	// 调用服务
	if err := server.Serve(lis); err != nil {
		log.Fatalf("fail to server: %v", err)
	}
}
