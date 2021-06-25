package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/zewei1022/lemon-gin-web-framework/global"
	"time"
)

var jwtSecret = []byte(global.LGWF_CONFIG.JWT.Secret)

type Claims struct {
	ID uint
	jwt.StandardClaims
}

func GenerateToken(id uint) (string, error) {
	expireTime := time.Now().Unix() + global.LGWF_CONFIG.JWT.Ttl

	claims := Claims{
		id,
		jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "lemon-gin-web-framework",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
	}

	return nil, err
}
