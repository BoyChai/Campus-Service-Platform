package middle

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

// SECRET 加解密因子
const (
	SECRET = "ikun"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//对登录接口放行
		if len(ctx.Request.URL.String()) >= 6 && ctx.Request.URL.String()[0:6] == "/login" {
			ctx.Next()
		} else {
			//获取Header中的Authorization
			token := ctx.Request.Header.Get("Authorization")
			if token == "" {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"msg":  "请求未携带token，无权限访问",
					"data": nil,
				})
				ctx.Abort()
				return
			}
			// parseToken 解析token包含的信息
			claims, err := parseToken(token)
			if err != nil {
				//token延期错误
				if err.Error() == "TokenExpired" {
					ctx.JSON(http.StatusBadRequest, gin.H{
						"msg":  "授权已过期",
						"data": nil,
					})
					ctx.Abort()
					return
				}
				//其他解析错误
				ctx.JSON(http.StatusBadRequest, gin.H{
					"msg":  err.Error(),
					"data": nil,
				})
				ctx.Abort()
				return
			}
			// 继续交由下一个路由处理,并将解析出的信息传递下去
			ctx.Set("claims", claims)
			ctx.Next()
		}
	}
}

// parseToken 解析token
func parseToken(tokenString string) (user map[string]interface{}, err error) {
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
