package util

import (
	jwt "github.com/dgrijalva/jwt-go"
	"memo-RPC/userserver/conf"
	"time"
)

var jwtSecret = []byte(conf.Cfg.App.JwtSecret)

type Claims struct {
	UID      uint   `json:"uid"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(uid uint, username string) (string, error) {

	// 设置 token 有效时间，3小时
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		uid,
		username,
		jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			// 指定 token 发行人信息
			Issuer:   "grpc-user-service",
			IssuedAt: time.Now().Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 该方法内部生成签名字符串，再用于获取完整、已签名的 token
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

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
