package main

import (
	"Campus-Service-Platform/controller"
	"Campus-Service-Platform/dao"
	"Campus-Service-Platform/middle"
	"Campus-Service-Platform/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化注册验证码
	utils.Code = make(map[string]string)
	// 初始化upyun对象
	utils.InitUpyun()
	// 连接数据库
	dao.Client()
	// 自动建表
	dao.AutoTables()
	// 创建gin
	r := gin.Default()
	// 加载中间件
	r.Use(middle.CORS())
	r.Use(middle.JWTAuth())
	// 加载路由
	controller.Router.InitApiRouter(r)
	// 运行WebService
	r.Run()
}
