package service

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/metadata"
	"log"
)

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
	} else {
		log.Printf("service getTokenFromContext MD: %v", md)
	}
	// md 的类型是 type MD map[string][]string
	token, ok := md["authorization"]
	if !ok || len(token) == 0 {
		return "", fmt.Errorf("ErrNoAuthorizationInMetadata")
	}
	// 因此，token 是一个字符串数组，我们只用了 token[0]
	return token[0], nil
}

func CheckAuthority(ctx context.Context) (username string, err error) {
	tokenStr, err := getTokenFromContext(ctx)
	if err != nil {
		log.Fatalf("Check Authority: get token from context error, %v", err)
		return "", err
	}
	log.Printf("Token content: %v", tokenStr)

	//token, err := jwt.ParseWithClaims(tokenStr, &clientClaims, func(token *jwt.Token) (interface{}, error) {
	//	if token.Header["alg"] != "HS256" {
	//		log.Fatalf("Check Authority: ErrInvalidAlgorithm, %v", err)
	//	}
	//	return []byte("very secret"), nil
	//})
	//if err != nil {
	//	log.Fatalf("Check Authority: jwt parse error, %v", err)
	//	return "", err
	//}

	//var clientClaims *Claims
	//clientClaims, err = ParseToken(tokenStr)
	//if time.Now().Unix() > clientClaims.ExpiresAt {
	//	log.Fatalf("Check Authority: ErrInvalidToken, %v", err)
	//	return "", err
	//} else if err != nil {
	//	log.Fatalf("Check Authority: jwt parse error, %v", err)
	//	return "", err
	//}
	//if !tokenStr.Valid {
	//	log.Fatalf("Check Authority: ErrInvalidToken, %v", err)
	//	return "", err
	//}

	return "Anxiu", nil
}

var jwtSecret = []byte("")

// ParseToken 根据传入的 token 值获取到 Claims 对象信息，进而获取其中的用户名和密码
func ParseToken(token string) (*Claims, error) {

	// 用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回 *Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// 从 tokenClaims 中获取到 Claims 对象，并使用断言，将该对象转换为我们自己定义的 Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
