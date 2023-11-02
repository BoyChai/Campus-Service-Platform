package db

import "time"
import "gorm.io/gorm"

// User 用户表
type User struct {
	// ID用户的唯一标识
	ID uint `gorm:"primaryKey;auto_increment"`
	// 用户名
	Name string
	// 电话号
	Number string
	// 密码(md5(Pass))
	Pass string
	// 图片地址
	ImgUrl string
	// 微信id
	WxID string
	// 角色
	Role int
}

// WorkOrder 工单表
type WorkOrder struct {
	ID         uint `gorm:"primaryKey;auto_increment"`
	CreateUser uint
	OrderType  int
	CreateImg  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	CompleteAT time.Time
	Info       string
}
