package service

import "github.com/golang-jwt/jwt/v4"

type JWTClaims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"user_name"`
	jwt.RegisteredClaims
}

type UserService struct {
}
