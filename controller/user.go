package controller

import (
	"Campus-Service-Platform/dao"
	"Campus-Service-Platform/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var User user

type user struct {
}

// Login 登录
func (u *user) Login(ctx *gin.Context) {
	// 参数绑定
	params := new(struct {
		Number string `form:"number" binding:"required"`
		Pass   string `form:"pass"  binding:"required"`
		WxID   string `form:"wx_id"`
	})
	if err := ctx.Bind(&params); err != nil {
		fmt.Println("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	check, err := dao.Dao.CheckUser(params.Number, utils.CalculateMD5Hash(params.Pass))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	token, err := utils.GenerateToken(check.GetID(), int(check.Role))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  fmt.Sprint("生成token出现错误:", err.Error()),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "登录成功!",
		"data": token,
	})
	return
}
