package util

import (
	"context"
	"log"
)

// TokenAuth 通过实现 gRPC 中默认定义的 PerRPCCredentials，提供用于自定义认证的接口，它的作用是将所需的安全认证信息添加到每个 RPC 方法的上下文中。
type TokenAuth struct {
	Token string
}

// GetRequestMetadata 获取当前请求认证所需的元数据
func (auth *TokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	log.Printf("Get RequestMetadata: Auth Token: %v", auth.Token)
	return map[string]string{"Authorization": auth.Token}, nil
}

// RequireTransportSecurity 是否需要基于 TLS 认证进行安全传输
func (auth *TokenAuth) RequireTransportSecurity() bool {
	return true
}
