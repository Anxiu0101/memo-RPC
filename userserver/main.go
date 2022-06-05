package main

import (
	"google.golang.org/grpc/credentials"
	"log"
	"memo-RPC/userserver/conf"
	"memo-RPC/userserver/ecommerce"
	"memo-RPC/userserver/model"
	"memo-RPC/userserver/service"
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
	//var interceptor grpc.UnaryServerInterceptor
	//interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	//	// 拦截普通方法请求，验证Token
	//	err = service.CheckAuthority(ctx)
	//	if err != nil {
	//		return
	//	}
	//	// 继续处理请求
	//	return handler(ctx, req)
	//}

	var opts []grpc.ServerOption

	// TLS 认证
	certs, err := credentials.NewServerTLSFromFile("../certs/server-key.pem", "../certs/server-req.csr")
	if err != nil {
		log.Printf("Failed to generate credentials %v", err)
	}

	opts = append(opts, grpc.Creds(certs))

	// 创建新 grpc 服务
	// TODO 分离获取令牌功能和用户信息操作功能
	//server := grpc.NewServer(grpc.UnaryInterceptor(interceptor))
	server := grpc.NewServer()

	// 将服务注册到服务端
	ecommerce.RegisterUserServiceServer(server, &service.UserService{})
	log.Printf("server listening at %v", lis.Addr())

	// 调用服务
	if err := server.Serve(lis); err != nil {
		log.Fatalf("fail to server: %v", err)
	}
}
