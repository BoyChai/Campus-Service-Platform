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
	Name string `gorm:"not null"`
	// 电话号
	Number string `gorm:"unique;not null"`
	// 密码(md5(Pass))
	Pass string `gorm:"not null"`
	// 图片地址
	ImgUrl string `gorm:"not null"`
	// 微信id
	WxID string `gorm:"unique"`
	// 角色
	Role Role `gorm:"not null"`
}

// WorkOrder 工单表
type WorkOrder struct {
	// ID 订单的唯一标识
	ID uint `gorm:"primaryKey;auto_increment"`
	// CreateUser 创建订单的用户id
	CreateUser uint `gorm:"not null"`
	// OrderType 订单类型
	OrderType OrderType `gorm:"not null"`
	// CreatedAt 创建时间
	CreatedAt time.Time
	// UpdatedAt 订单更新时间
	UpdatedAt time.Time
	// DeletedAt 订单删除时间
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// CompleteAT 订单完成时间
	CompleteAT time.Time
	// 订单创建时的一些初始信息
	Info JSON `gorm:"type:json;not null"`
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
	// Work 工作者
	Work
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
