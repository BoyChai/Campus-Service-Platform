package main

import (
	"Campus-Service-Platform/dao"
	"Campus-Service-Platform/middle"
	"github.com/gin-gonic/gin"
)

func main() {
	// 连接数据库
	dao.Client()
	// 自动建表
	dao.AutoTables()
	// 创建gin
	r := gin.Default()
	// 加载中间件
	r.Use(middle.CORS())
	r.Use(middle.JWTAuth())
	// 运行WebService
	r.Run()
}
