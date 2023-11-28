package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
	"wow-admin/config"
)

var (
	ErrTokenExpired     = errors.New("token 已过期, 请重新登录")
	ErrTokenNotValidYet = errors.New("token 无效, 请重新登录")
	ErrTokenMalformed   = errors.New("token 不正确, 请重新登录")
	ErrTokenInvalid     = errors.New("这不是一个 token, 请重新登录")
)

type MyClaims struct {
	UserId int    `json:"userId"`
	Role   string `json:"role"`
	UUID   string `json:"uuid"`
	jwt.RegisteredClaims
}

type MyJWT struct {
	// secret 签名
	Secret []byte
}

func GetJWT() *MyJWT {
	// secret 签名
	return &MyJWT{[]byte(config.Cfg.JWT.Secret)}
}

/*
*
生成toekn
*/
func (j *MyJWT) GenToken(userId int, role string, uuid string) (string, error) {
	claims := MyClaims{
		UserId: userId,
		Role:   role,
		UUID:   uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.Cfg.JWT.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.Cfg.JWT.Expire) * time.Hour)),
		},
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用指定的 secret 签名并获得完整编码后的字符串 token
	return token.SignedString(j.Secret)
}

func (j *MyJWT) ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.Secret, nil
	})

	if err != nil {
		if vError, ok := err.(*jwt.ValidationError); ok {
			if vError.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if vError.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
			} else if vError.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}

		}
	}

	// 校验token
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}
