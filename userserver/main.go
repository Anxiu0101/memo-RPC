package main

import (
	"google.golang.org/grpc/credentials"
	"log"
	"memo-RPC/userserver/conf"
	pb "memo-RPC/userserver/ecommerce"
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

	// 监听本地端口
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	var opts []grpc.ServerOption

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

	// 使用证书文件和密钥文件为服务端构造 TLS 凭证
	// FIXME 使用 SAN 代替 x509，或者将 go 版本降级至 1.15 以下
	certs, err := credentials.NewServerTLSFromFile("../certs/server.pem", "../certs/server.key")
	if err != nil {
		log.Printf("Failed to generate credentials %v", err)
	} else {
		opts = append(opts, grpc.Creds(certs))
	}

	// 创建新 grpc 服务
	// TODO 分离获取令牌功能和用户信息操作功能
	server := grpc.NewServer(opts...)

	// 将服务注册到服务端
	pb.RegisterUserServiceServer(server, &service.UserService{})
	log.Printf("User server listening at %v", lis.Addr())

	// 调用服务
	if err := server.Serve(lis); err != nil {
		log.Fatalf("fail to server: %v", err)
	}
}
