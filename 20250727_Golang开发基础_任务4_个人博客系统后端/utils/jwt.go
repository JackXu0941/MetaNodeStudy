package utils

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID             uint   `json:"user_id"`  // 用户ID
	Username           string `json:"username"` // 用户名
	jwt.StandardClaims        // 标准声明（包含过期时间等）
}

var jwtSecret = []byte("Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTM2OTQ0NjEsImlkIjo5LCJ1c2VybmFtZSI6IjIyMjIifQ.p1MXQdvy2lA5lhyi9T6aNbomGKPnvCO5ZDZ-2mIzwTM")

// ParseToken 解析JWT token
func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token无效")
	}

	return claims, nil
}
