package service

import (
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
	"wow-admin/model"
)

var singedKey = []byte("1234567890!@#$%^&*()_+-=")

//生成token
func GenerateToken(userId int, unionid, phone, loginType string) (*model.UserTokenClaim, string) {
	userTkClaim := &model.UserTokenClaim{
		UserId:    userId,
		UnionId:   unionid,
		LoginType: loginType,
		Token:     hex.EncodeToString(uuid.NewV4().Bytes()),
		Timestamp: time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userTkClaim)
	ss, _ := token.SignedString(singedKey)
	return userTkClaim, ss
}

//解码token
func DecodeToken(tokenString string) (*model.UserTokenClaim, error) {
	if tokenString == "" {
		return nil, errors.New("No token ")
	}
	token, err := jwt.ParseWithClaims(tokenString, &model.UserTokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		return singedKey, nil
	})

	if err != nil {
		return nil, err
	}

	if token == nil {
		return nil, errors.New("Error token ")
	}

	claims, ok := token.Claims.(*model.UserTokenClaim)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("%v %v", claims.Token, claims.StandardClaims.ExpiresAt)
	}
	userService := UserBasicService{}
	_, err = userService.GetWeChatUser(claims.UserId)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
