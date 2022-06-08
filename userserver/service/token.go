package service

import (
	"github.com/dgrijalva/jwt-go"
	"memo-RPC/userserver/conf"
	"time"
)

func CreateToken(username string) (token string, err error) {

	claims := Claims{
		jwt.StandardClaims{
			Issuer:    "admin",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(3 * time.Hour).Unix(),
		},
		username,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var jwtSecret = []byte(conf.Cfg.App.JwtSecret)
	return tokenClaims.SignedString(jwtSecret)
}

// Claims defines the struct containing the token claims.
type Claims struct {
	jwt.StandardClaims

	// Username defines the identity of the user.
	Username string `json:"username"`
}
