package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// CheckAuthority 验证token
func CheckAuthority(ctx context.Context) error {

	// Get metadata from Context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "获取Token失败")
	}
	var (
		token string
	)
	if value, ok := md["token"]; ok {
		token = value[0]
	}
	if token == "0" {
		return status.Errorf(codes.Unauthenticated, "Token无效: token=%s", token)
	}
	return nil
}
