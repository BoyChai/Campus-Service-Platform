package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Client 数据库链接
func Client() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"User",
		"Pass",
		"Host",
		"Port",
		"Database",
	)
	Dao.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 逻辑外键（代码里自动体验外键关系）
	})
	if err != nil {
		fmt.Println("链接数据库错误:", err)
		return
	}
}

func AutoTables() {
	var err error
	// 用户表
	err = Dao.db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err)
	}
	// 工单表
	err = Dao.db.AutoMigrate(&WorkOrder{})
	if err != nil {
		fmt.Println(err)
	}
	// 聊天表
	err = Dao.db.AutoMigrate(&Chat{})
	if err != nil {
		fmt.Println(err)
	}
}
