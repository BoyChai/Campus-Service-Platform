package controller

import "github.com/gin-gonic/gin"

// Router 实例化router类型对象，首字母大写用于跨包调用
var Router router

// 声明router结构体
type router struct{}

func (r *router) InitApiRouter(router *gin.Engine) {
	router.GET("/SendSMS", User.SendSMS).
		GET("/Signup", User.Signup).
		GET("/login", User.Login)

}
