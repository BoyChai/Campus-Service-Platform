package middle

import (
	"Campus-Service-Platform/utils"
	"github.com/gin-gonic/gin"
	"net/http"
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
			claims, err := utils.ParseToken(token)
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
