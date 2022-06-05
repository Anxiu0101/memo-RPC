package service

//import (
//	"context"
//	"fmt"
//	jwt "github.com/dgrijalva/jwt-go"
//	"google.golang.org/grpc/codes"
//	"google.golang.org/grpc/metadata"
//	"google.golang.org/grpc/status"
//	"memo-RPC/userserver/conf"
//	"memo-RPC/userserver/model"
//	"time"
//)
//
//// CheckAuthority 验证token
//func CheckAuthority(ctx context.Context) error {
//
//	// Get metadata from Context
//	md, ok := metadata.FromIncomingContext(ctx)
//	if !ok {
//		return status.Errorf(codes.Unauthenticated, "获取Token失败")
//	}
//	var (
//		appID     string
//		appSecret string
//	)
//	if value, ok := md["app_id"]; ok {
//		appID = value[0]
//	}
//	if value, ok := md["app_secret"]; ok {
//		appSecret = value[0]
//	}
//	var user model.User
//	model.DB.Where("username = ?", appID).Find(&user)
//	if appID != "grpc_token" || appSecret != "123456" {
//		return status.Errorf(codes.Unauthenticated, "Token无效: app_id=%s, app_secret=%s", appID, appSecret)
//	}
//	return nil
//}
//
//var jwtSecret = []byte(conf.Cfg.App.JwtSecret)
//
//type Claims struct {
//	ID        uint   `json:"id"`
//	Username  string `json:"username"`
//	jwt.StandardClaims
//}
//
//func CreateToken(uid uint, username string) (tokenString string) {
//
//	accessExpireTime := time.Now().Add(3 * time.Hour)
//
//	claims := Claims{
//		ID: uid,
//		Username: username,
//		jwt.StandardClaims{Issuer: "admin"},
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims})
//	tokenString, err := token.SignedString([]byte("verysecret"))
//	if err != nil {
//		panic(err)
//	}
//	return tokenString
//}

import (
	"context"
	"fmt"
	"memo-RPC/userserver/conf"
	"time"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/metadata"
)

func CreateToken(username string) (token string, err error) {
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":      "memo-RPC-admin",
		"aud":      "memo-RPC-admin",
		"nbf":      time.Now().Unix(),
		"exp":      time.Now().Add(3 * time.Hour).Unix(),
		"sub":      "user",
		"username": username,
	})
	println("conf.Cfg.App.JwtSecret: ", conf.Cfg.App.JwtSecret)
	var jwtSecret = []byte(conf.Cfg.App.JwtSecret)
	return tokenClaims.SignedString(jwtSecret)
}

// AuthToken 自定义认证
type AuthToken struct {
	Token string
}

func (c AuthToken) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": c.Token,
	}, nil
}

func (c AuthToken) RequireTransportSecurity() bool {
	return false
}

// Claims defines the struct containing the token claims.
type Claims struct {
	jwt.StandardClaims

	// Username defines the identity of the user.
	Username string `json:"username"`
}

// Step1. 从 context 的 metadata 中，取出 token

func getTokenFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("ErrNoMetadataInContext")
	}
	// md 的类型是 type MD map[string][]string
	token, ok := md["authorization"]
	if !ok || len(token) == 0 {
		return "", fmt.Errorf("ErrNoAuthorizationInMetadata")
	}
	// 因此，token 是一个字符串数组，我们只用了 token[0]
	return token[0], nil
}

func CheckAuth(ctx context.Context) (username string) {
	tokenStr, err := getTokenFromContext(ctx)
	if err != nil {
		panic("get token from context error")
	}
	var clientClaims Claims
	token, err := jwt.ParseWithClaims(tokenStr, &clientClaims, func(token *jwt.Token) (interface{}, error) {
		if token.Header["alg"] != "HS256" {
			panic("ErrInvalidAlgorithm")
		}
		return []byte("verysecret"), nil
	})
	if err != nil {
		panic("jwt parse error")
	}

	if !token.Valid {
		panic("ErrInvalidToken")
	}

	return clientClaims.Username
}
