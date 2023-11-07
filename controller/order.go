package controller

import (
	"Campus-Service-Platform/dao"
	"Campus-Service-Platform/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

var Order order

type order struct {
}

// GetOrders 获取订单
func (o *order) GetOrders(ctx *gin.Context) {
	params := new(struct {
		Status string `form:"status" binding:"required"`
	})
	if err := ctx.Bind(&params); err != nil {
		fmt.Println("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	intValue, err := strconv.Atoi(params.Status)
	if err != nil {
		fmt.Println("传入订单类型错误:" + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "传入订单类型错误:" + err.Error(),
			"data": "",
		})
		return
	}

	orders, err := dao.Dao.QueryStatusOrder(dao.Role(intValue))
	if err != nil {
		fmt.Println("通过状态查询订单失败:" + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": "",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取成功!",
		"data": orders,
	})
}

// CreateOrder 创建订单
func (o *order) CreateOrder(ctx *gin.Context) {
	// 拿到身份
	claims, _ := ctx.Get("claims")
	id := claims.(map[string]interface{})["id"]
	//参数绑定
	params := new(struct {
		OrderType int    `form:"type" binding:"required"`
		Info      string `form:"info" binding:"required"`
	})
	if err := ctx.Bind(&params); err != nil {
		fmt.Println("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	file, err := ctx.FormFile("img")
	if err != nil {
		fmt.Println("img请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	img, err := file.Open()
	defer img.Close()
	if err != nil {
		fmt.Println("img文件Open失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	fileName := fmt.Sprint(utils.GetNumber() + "." + strings.Split(file.Filename, ".")[1])
	url, err := utils.PutOrderImg(fileName, img)
	if err != nil {
		fmt.Println("img文件上传失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	order, err := dao.Dao.CreateOrder(utils.GetUint(fmt.Sprint(id)), dao.OrderType(params.OrderType), dao.JSON{
		"img":  url,
		"info": params.Info,
	})
	if err != nil {
		fmt.Println("订单创建失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "订单创建成功",
		"data": order,
	})
	return
}

func (o *order) ReceivingOrder(ctx *gin.Context) {
	// 拿到身份
	claims, _ := ctx.Get("claims")
	id := claims.(map[string]interface{})["id"]
	//参数绑定
	params := new(struct {
		OrderID string `form:"id" binding:"required"`
	})
	if err := ctx.Bind(&params); err != nil {
		fmt.Println("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	err := dao.Dao.ReceivingOrder(params.OrderID, utils.GetUint(fmt.Sprint(id)))
	if err != nil {
		fmt.Println("接单失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "接单成功!",
		"data": nil,
	})
}
