package dao

import (
	"gorm.io/gorm"
	"time"
)

var Dao dao

type dao struct {
	db *gorm.DB
}
type JSON map[string]interface{}

// User 用户表
type User struct {
	// ID用户的唯一标识
	ID uint `gorm:"primaryKey;auto_increment"`
	// 用户名
	Name string
	// 电话号
	Number string `gorm:"unique"`
	// 密码(md5(Pass))
	Pass string
	// 图片地址
	ImgUrl string
	// 微信id
	WxID string `gorm:"unique"`
	// 角色
	Role Role
}

// WorkOrder 工单表
type WorkOrder struct {
	ID         uint `gorm:"primaryKey;auto_increment"`
	CreateUser uint
	OrderType  OrderType
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	CompleteAT time.Time
	Info       JSON `gorm:"type:json"`
}

// Role 角色类别
type Role int

// OrderType 订单类型
type OrderType int

const (
	// Admin 超级管理员
	Admin Role = iota
	// Leader 领导
	Leader
	// Ordinary 普通用户
	Ordinary
)

const (
	// Power 电力故障
	Power OrderType = iota
	// Network 网络故障
	Network
	// Water 水源故障
	Water
	// HVAC 暖通空调故障
	HVAC
	// Device 设备故障
	Device
	// Construction 建筑设施
	Construction
	// SecuritySystem 安全系统
	SecuritySystem
	// CampusTransportation 校园交通
	CampusTransportation
	// Health 卫生
	Health
)
