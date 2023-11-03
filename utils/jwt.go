package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// SECRET 加解密因子
const (
	SECRET = "ikun"
)

// GenerateToken 生成token
func GenerateToken(id string, role int) (string, error) {
	// 创建的token有效期是一周
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   id,
		"role": role,
		"exp":  time.Now().Add(time.Hour * 168).Unix(),
	})
	return token.SignedString([]byte(SECRET))
}

// ParseToken 解析token
func ParseToken(tokenString string) (user map[string]interface{}, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 校验签名是否被篡改
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("not authorization")
		}
		//返回密钥与上面签发时保持一致
		return []byte(SECRET), nil
	})
	if err != nil {
		fmt.Println("parse token failed ", err)
		//处理token解析后的各种错误
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, errors.New("TokenMalformed")
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("TokenExpired")
		} else if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, errors.New("TokenNotValidYet")
		} else {
			return nil, errors.New("TokenInvalid")
		}
	}
	return token.Claims.(jwt.MapClaims), nil
}
