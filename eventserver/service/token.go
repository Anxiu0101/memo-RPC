package service

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

// Claims defines the struct containing the token claims.
type Claims struct {
	jwt.StandardClaims

	// Username defines the identity of the user.
	Username string `json:"username"`
}

func CheckAuthority(ctx context.Context) (username string, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Fatalln("Check Authority: ErrNoMetadataInContext")
		return "", fmt.Errorf("ErrNoMetadataInContext")
	}
	// md 的类型是 type MD map[string][]string
	// Attention: metadata 里的 key 将会全部被转化为小写的，只能用小写的查
	token, ok := md["authorization"]
	if !ok || len(token) == 0 {
		log.Fatalln("Check Authority: ErrNoAuthorizationInMetadata")
		return "", fmt.Errorf("ErrNoAuthorizationInMetadata")
	}

	var tokenStr = token[0]

	var clientClaims *Claims
	clientClaims, err = ParseToken(tokenStr)

	if err != nil {
		log.Fatalf("Check Authority: jwt parse error, %v", err)
		return "", err
	} else if time.Now().Unix() > clientClaims.ExpiresAt {
		log.Fatalf("Check Authority: ErrInvalidToken, %v", err)
		return "", err
	}

	return clientClaims.Username, nil
}

var jwtSecret = []byte("23347$040412")

// ParseToken 根据传入的 token 值获取到 Claims 对象信息，进而获取其中的用户名和密码
func ParseToken(token string) (*Claims, error) {

	// 用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回 *Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		log.Printf("JWT Secret: %v", jwtSecret)
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
