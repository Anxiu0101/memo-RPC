package service

import (
	"context"
	"log"
	pb "memo-RPC/userserver/ecommerce"
	"memo-RPC/userserver/model"
)

type UserService struct {
}

func (*UserService) Login(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {

	// 查询用户是否存在
	// 错误情况：用户被封禁，用户不存在，数据库错误
	var user model.User
	if err := model.DB.Where("username = ?", req.Username).Find(&user).Error; err != nil {
		return &pb.UserLoginResponse{
			Token: "User not exits",
		}, err
	}

	// 验证用户密码是否正确
	if err := user.CheckPassword(req.Password); err != nil {
		return &pb.UserLoginResponse{
			Token: "User not exits",
		}, err
	}

	// 生成 token
	token, err := CreateToken(req.Username)
	if err != nil {
		log.Println(err)
		return &pb.UserLoginResponse{
			Token: "Fail to Create Token",
		}, err
	}

	return &pb.UserLoginResponse{
		Token: token,
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
