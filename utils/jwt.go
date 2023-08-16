package utils

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/weilaim/wmsystem/config"
)

// 定义 token 相关 error
var (
	ErrTokenExpired     = errors.New("token 已过期, 请重新登录")
	ErrTokenNotValidYet = errors.New("token 无效, 请重新登录")
	ErrTokenMalformed   = errors.New("token 不正确, 请重新登录")
	ErrTokenInvalid     = errors.New("这不是一个 token, 请重新登录")
)

// JWT工具类

type MyClaims struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
	UUID   string `json:"uuid"`
	jwt.RegisteredClaims
}

type MyJWt struct {
	Secret []byte
}

func GetJWT() *MyJWt {
	return &MyJWt{[]byte(config.Cfg.JWT.Secret)}
}

func (j *MyJWt) ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
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

	// 校验 token
	if caims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return caims, nil
	}

	return nil, ErrTokenInvalid
}
