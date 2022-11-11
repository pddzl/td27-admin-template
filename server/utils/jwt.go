package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"server/model/system/request"
)

type JWT struct {
	SigningKey []byte
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}
