package dao

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"time"
)

var Dao dao

type dao struct {
	db *gorm.DB
}
type JSON map[string]interface{}

func (j JSON) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan is used to convert a database value to the custom JSON type.
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Invalid JSON format")
	}
	return json.Unmarshal(bytes, j)
}

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
	//  OrderStatus 订单状态
	OrderStatus OrderStatus `gorm:"not null"`
	// CreatedAt 创建时间
	CreatedAt time.Time
	// UpdatedAt 订单更新时间
	UpdatedAt time.Time
	// DeletedAt 订单删除时间
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// CompleteAT 订单完成时间
	CompleteAT *time.Time
	// 订单创建时的一些初始信息
	Info JSON `gorm:"type:json;not null"`
}

// Chat 聊天表
type Chat struct {
	// 消息id
	ID uint `gorm:"primaryKey;auto_increment"`
	// 消息订单
	OrderID uint `gorm:"not null"`
	// 发送角色
	Sender Role `gorm:"not null"`
	// 发送时间
	SendTime time.Time `gorm:"not null"`
	// 消息内容
	Message JSON `gorm:"type:json;not null"`
}

// Role 角色类别
type Role int

// OrderType 订单类型
type OrderType int

// OrderStatus 订单状态
type OrderStatus int

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

const (
	// Pending 待处理
	Pending OrderStatus = iota
	// InProgress 处理中
	InProgress
	// WaitingConfirm 等待确认
	WaitingConfirm
	// Success 处理成功
	Success
)
