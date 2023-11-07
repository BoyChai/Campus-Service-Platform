package controller

import (
	"Campus-Service-Platform/dao"
	"Campus-Service-Platform/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Chat 实例化chat类型对象，首字母大写用于跨包调用
var Chat chat

// 声明router结构体
type chat struct{}

func (c *chat) SendMessage(ctx *gin.Context) {
	// 拿到身份
	claims, _ := ctx.Get("claims")
	role := claims.(map[string]interface{})["role"]
	params := new(struct {
		OrderID string   `form:"id" binding:"required"`
		Message dao.JSON `form:"msg" binding:"required"`
	})
	if err := ctx.Bind(&params); err != nil {
		fmt.Println("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	err := dao.Dao.SendMessage(utils.GetUint(params.OrderID), role.(dao.Role), params.Message)
	if err != nil {
		fmt.Println("发送消息, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "发送成功！",
		"data": nil,
	})
}

func (c *chat) GetMsg(ctx *gin.Context) {
	params := new(struct {
		OrderID uint `form:"id" binding:"required"`
	})
	if err := ctx.Bind(&params); err != nil {
		fmt.Println("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	msg, err := dao.Dao.QueryOrderMsg(params.OrderID)
	if err != nil {
		fmt.Println("查询订单消息错误: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"msg":  "查询成功",
		"data": msg,
	})
	return
}
