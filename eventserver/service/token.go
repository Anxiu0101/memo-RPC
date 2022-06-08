package service

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/metadata"
	"log"
)

// CheckAuthority 验证token
//func CheckAuthority(ctx context.Context) error {
//
//	log.Println("Checking Token")
//
//	// Get metadata from Context
//	md, ok := metadata.FromIncomingContext(ctx)
//	if !ok {
//		return status.Errorf(codes.Unauthenticated, "获取Token失败")
//	}
//	var (
//		token string
//	)
//	if value, ok := md["token"]; ok {
//		token = value[0]
//	}
//	if token == "0" {
//		return status.Errorf(codes.Unauthenticated, "Token无效: token=%s", token)
//	}
//	return nil
//}

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
	token, ok := md["Authorization"]
	if !ok || len(token) == 0 {
		return "", fmt.Errorf("ErrNoAuthorizationInMetadata")
	}
	// 因此，token 是一个字符串数组，我们只用了 token[0]
	return token[0], nil
}

func CheckAuthority(ctx context.Context) (username string, err error) {
	tokenStr, err := getTokenFromContext(ctx)
	log.Printf("tokenStr Content: %v", tokenStr)
	if err != nil {
		log.Fatalf("Check Authority: get token from context error, %v", err)
		return "", err
	}
	var clientClaims Claims
	token, err := jwt.ParseWithClaims(tokenStr, &clientClaims, func(token *jwt.Token) (interface{}, error) {
		if token.Header["alg"] != "HS256" {
			log.Fatalf("Check Authority: ErrInvalidAlgorithm, %v", err)
		}
		return []byte("very secret"), nil
	})
	if err != nil {
		log.Fatalf("Check Authority: jwt parse error, %v", err)
		return "", err
	}

	if !token.Valid {
		log.Fatalf("Check Authority: ErrInvalidToken, %v", err)
		return "", err
	}

	return clientClaims.Username, nil
}
