package model

import "github.com/golang-jwt/jwt/v4"

type UserTokenClaim struct {
	Token     string `json:"token"`
	UserId    int    `json:"userId"`
	UnionId   string `json:"unionid"`
	LoginType string `json:"loginType"`
	Phone     string `json:"phone"`
	Timestamp int64  `json:"timestamp"`

	jwt.StandardClaims
}
