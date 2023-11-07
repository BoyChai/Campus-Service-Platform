package controller

import "github.com/gin-gonic/gin"

// Router 实例化router类型对象，首字母大写用于跨包调用
var Router router

// 声明router结构体
type router struct{}

func (r *router) InitApiRouter(router *gin.Engine) {
	router.
		// 用户相关
		POST("/user/sendsms", User.SendSMS).
		POST("/user/signup", User.Signup).
		POST("/user/login", User.Login).
		// 订单相关
		GET("/order/list", Order.GetOrders).
		POST("/order/create", Order.CreateOrder).
		POST("/order/receiving", Order.ReceivingOrder)
}
