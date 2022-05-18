package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "memo-RPC/proto/pb_go"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const PORT = "9001"

func main() {

	// 监听端口
	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	// 创建新 grpc 服务
	server := grpc.NewServer()

	// 将 Search 服务注册到服务端
	pb.RegisterSearchServiceServer(server, &SearchService{})
	log.Printf("server listening at %v", lis.Addr())

	// 调用服务
	if err := server.Serve(lis); err != nil {
		log.Fatalf("fail to server: %v", err)
	}
}
