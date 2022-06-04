package service

import (
	"context"
	pb "memo-RPC/userserver/ecommerce"
)

type UserService struct {
}

func (*UserService) Login(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	return &pb.UserLoginResponse{
		Token: "",
	}, nil
}

func (*UserService) Register(ctx context.Context, req *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
	return &pb.UserRegisterResponse{
		Id: 0,
	}, nil
}

func (*UserService) ShowUserInfo(ctx context.Context, req *pb.ShowUserInfoRequest) (*pb.ShowUserInfoResponse, error) {
	return &pb.ShowUserInfoResponse{
		Item: nil,
	}, nil
}
